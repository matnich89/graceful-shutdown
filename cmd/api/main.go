package main

import (
	"github.com/go-chi/chi/v5"
	"log"
)

func main() {
	router := chi.NewRouter()

	app := newApp(router)

	app.routes()

	log.Fatalln(app.serve())
}
