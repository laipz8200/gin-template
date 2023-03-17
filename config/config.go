package config

import (
	"encoding/json"
	"os"
)

var c config = config{}

func Init(filepath string) {
	byteValue, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(byteValue, &c); err != nil {
		panic(err)
	}
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
