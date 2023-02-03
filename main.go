package main

import (
	"_template_/api"
	"_template_/registry"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Connect to database
	config := gorm.Config{}
	config.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		},
	)
	db, err := gorm.Open(sqlite.Open(":memory:"), &config)
	if err != nil {
		log.Fatal(err)
	}

	// Init registry
	r := NewRegistry(db)
	registry.Init(r)

	if err := api.Run("127.0.0.1:8000"); err != nil {
		log.Fatal(err)
	}
}
