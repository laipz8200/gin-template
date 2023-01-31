package main

import (
	"_template_/internal/web"
	"log"
)

func main() {
	if err := web.Run("127.0.0.1:8000"); err != nil {
		log.Fatal(err)
	}
}
