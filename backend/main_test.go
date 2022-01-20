package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
	"github.com/yinkar/tototodo/backend/src"
)

type Api struct {
	rr *httptest.ResponseRecorder
}

func (a *Api) makeRequest(method, endpoint string) (err error) {

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return
	}

	var handler http.HandlerFunc

	switch endpoint {
	case "/version":
		handler = http.HandlerFunc(src.Version)
	case "/health":
		handler = http.HandlerFunc(src.Health)
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

	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	if err = json.Unmarshal(a.rr.Body.Bytes(), &actual); err != nil {
		return
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("The JSON body should be %v, but %v given", expected, actual)
	}

	return nil
}

func (a *Api) resetResponse(*godog.Scenario) {
	a.rr = httptest.NewRecorder()
}

func InitializeScenario(s *godog.ScenarioContext) {
	api := &Api{}

	s.BeforeScenario(api.resetResponse)

	s.Step(`^client makes a "(GET|POST)" request to "(\/[^"]*)"$`, api.makeRequest)
	s.Step(`^response code should be (\d+)$`, api.responseCodeShouldBe)
	s.Step(`^response body should match with:$`, api.responseShouldMatchWith)
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
