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
	mux.Use(app.Session.LoadAndSave)

	mux.Route("/", func(r chi.Router) {
		mux.Get("/", app.Home)
		mux.Post("/login", app.Login)

		mux.With(app.auth)
		mux.Get("/user/profile", app.Profile)

	})

	// static files
	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// mux satisfies the http.Handler interface
	return mux
}
