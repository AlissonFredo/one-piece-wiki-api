package main

import (
	"log"

	"github.com/AlissonFredo/one-piece-wiki-api/internal/app"
)

func main() {
	r := app.NewRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
