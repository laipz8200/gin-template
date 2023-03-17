package config

import (
	"encoding/json"
	"os"
)

var c config = config{}

func LoadFromFile(filepath string) config {
	byteValue, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var config config
	if err := json.Unmarshal(byteValue, &config); err != nil {
		panic(err)
	}
	return config
}

func SetConfig(config config) {
	c = config
}

type config struct {
	AppName string `json:"app_name"`
	Debug   bool   `json:"debug"`
	Secret  string `json:"secret"`
}

func AppName() string {
	return c.AppName
}

func Debug() bool {
	return c.Debug
}

func Secret() string {
	return c.Secret
}
