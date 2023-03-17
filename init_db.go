package main

import (
	"_template_/config"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
