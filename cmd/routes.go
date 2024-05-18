package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)
	mux.Get("/loged", OauthGoogleCallback)
	mux.Get("/login", OauthGoogleLogin)
	// static assets

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
