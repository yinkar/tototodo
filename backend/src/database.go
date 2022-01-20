package src

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	c "github.com/yinkar/tototodo/backend/_config"
)

func Connect() *gorm.DB {
	config := c.GetConfig()

	db, err := gorm.Open(config.DB.Driver,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
			config.DB.Username,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
			config.DB.Database,
			config.DB.Charset))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
