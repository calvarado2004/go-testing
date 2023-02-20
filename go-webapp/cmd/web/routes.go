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

	// register routes

	// static files

	// mux satisfies the http.Handler interface
	return mux
}
