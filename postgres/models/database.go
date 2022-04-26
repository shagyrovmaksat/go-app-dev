package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	connStr := "host = 127.0.0.1 port = 5432 user = postgres dbname = courses password = 09082002 sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic("Can't connect to database")
	}

	db.AutoMigrate(&Course{})
	DB = db
}
