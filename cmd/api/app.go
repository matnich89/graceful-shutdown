package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type app struct {
	router *chi.Mux
}

func newApp(router *chi.Mux) *app {
	return &app{router: router}
}

func (a *app) routes() {
	// Define your routes
	a.router.Get("/greeting", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello there"))
	})
}
