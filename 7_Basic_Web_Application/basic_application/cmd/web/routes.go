package main

import (
	"basicWebApp/pkg/config"
	"basicWebApp/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(a *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
