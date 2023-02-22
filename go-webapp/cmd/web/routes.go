package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)

	mux.Use(app.addIPToContext)

	// register routes
	mux.Get("/", app.Home)

	// static files
	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// mux satisfies the http.Handler interface
	return mux
}