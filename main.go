package main

import (
	"_template_/api"
	"_template_/config"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	c := config.LoadFromFile("config/config.json")
	config.SetConfig(c)

	// Connect to database
	dbConfig := dbConfig()
	gorm.Open(sqlite.Open(":memory:"), &dbConfig)

	server := api.NewServer()
	if err := server.Run("0.0.0.0:8000"); err != nil {
		log.Fatal(err)
	}
}
