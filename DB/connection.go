package DB

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost port=5445 user=mejiadev password=mejiadev dbname=biblioteca sslmode=disable TimeZone=AMerica/El_Salvador"
var DB *gorm.DB

func Connect() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal("Error: ", error)
	} else {
		log.Println("Database connected")
	}
}
