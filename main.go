package main

import (
	"_template_/api"
	"_template_/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	config.Init("config/config.json")
	api.Init()
}

func main() {
	// Connect to database
	dbConfig := dbConfig()

	gorm.Open(sqlite.Open(":memory:"), &dbConfig)

	if err := api.Run("0.0.0.0:8000"); err != nil {
		log.Fatal(err)
	}
}

func dbConfig() gorm.Config {
	c := gorm.Config{}
	if config.Debug() {
		c.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             0,
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
				LogLevel:                  logger.Info,
			},
		)
	} else {
		c.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             10 * time.Second,
				Colorful:                  true,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  logger.Warn,
			},
		)
	}
	return c
}
