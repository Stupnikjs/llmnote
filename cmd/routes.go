package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)

	mux.Get("/signup", app.RenderSignup)
	mux.Get("/error", app.RenderError)

	// static assets

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
