package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
	c "github.com/yinkar/tototodo/backend/_config"
	"github.com/yinkar/tototodo/backend/src"
)

var config = c.GetConfig()

func init() {

}

type Api struct {
	rr *httptest.ResponseRecorder
	db *sql.DB
}

var apiHandlers = &src.ApiHandlers{Db: src.Connect()}

func (a *Api) makeRequest(method, endpoint string) (err error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return
	}

	var handler http.HandlerFunc

	switch endpoint {
	case "/version":
		handler = http.HandlerFunc(apiHandlers.Version)
	case "/health":
		handler = http.HandlerFunc(apiHandlers.Health)
	case "/todos":
		if method == "GET" {
			handler = http.HandlerFunc(apiHandlers.GetTodos)
		} else if method == "POST" {
			handler = http.HandlerFunc(apiHandlers.CreateTodo)
		}
	default:
		return fmt.Errorf("Unknown endpoint %s", endpoint)
	}

	handler.ServeHTTP(a.rr, req)

	return
}

func (a *Api) responseCodeShouldBe(code int) error {
	if code != a.rr.Code {
		return fmt.Errorf("The response code should be %v, but %v given", code, a.rr.Code)
	}
	return nil
}

func (a *Api) responseShouldMatchWith(body *godog.DocString) (err error) {
	var expected, actual interface{}

	err = json.Unmarshal([]byte(body.Content), &expected)
	if err != nil {
		return
	}

	err = json.Unmarshal(a.rr.Body.Bytes(), &actual)
	if err != nil {
		return
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("The JSON body should be %v, but %v given", expected, actual)
	}

	return nil
}

func (a *Api) responseLookLikeThis(body *godog.DocString) (err error) {
	expected := make([]map[string]interface{}, 0)
	actual := make([]map[string]interface{}, 0)

	err = json.Unmarshal([]byte(body.Content), &expected)
	if err != nil {
		return
	}

	err = json.Unmarshal(a.rr.Body.Bytes(), &actual)
	if err != nil {
		return
	}

	for i := range actual {
		delete(actual[i], "Id")
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("The JSON body should be %v, but %v given", expected, actual)
	}

	return nil
}

func (a *Api) thereAreTodos(todos *godog.Table) error {
	for i := 1; i < len(todos.Rows); i++ {
		query := "INSERT INTO todos(content) VALUES ('%s')"

		insert, err := a.db.Query(fmt.Sprintf(query, todos.Rows[i].Cells[0].Value))

		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
	}
	return nil
}

func (a *Api) resetResponse(*godog.Scenario) {
	a.rr = httptest.NewRecorder()

	if a.db != nil {
		a.db.Close()
	}
	db, err := sql.Open(config.DB.Driver,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
			config.DB.Username,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
			config.DB.Database,
			config.DB.Charset))
	if err != nil {
		panic(err)
	}

	a.db = db
}

func InitializeScenario(s *godog.ScenarioContext) {
	api := &Api{}

	s.BeforeScenario(api.resetResponse)

	s.Step(`^client makes a "(GET|POST)" request to "(\/[^"]*)"$`, api.makeRequest)
	s.Step(`^response code should be (\d+)$`, api.responseCodeShouldBe)
	s.Step(`^response body should match with:$`, api.responseShouldMatchWith)
	s.Step(`^there are todos:$`, api.thereAreTodos)
	s.Step(`^response body should look like:$`, api.responseLookLikeThis)
}

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                "tototodo",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	os.Exit(status)
}
