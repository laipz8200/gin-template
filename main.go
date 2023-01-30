package main

import (
	"log"
	"template/internal/web"
)

func main() {
	if err := web.Run("127.0.0.1:8000"); err != nil {
		log.Fatal(err)
	}
}
