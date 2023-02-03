package main

import (
	"_template_/api"
	"log"
)

func main() {
	if err := api.Run("127.0.0.1:8000"); err != nil {
		log.Fatal(err)
	}
}
