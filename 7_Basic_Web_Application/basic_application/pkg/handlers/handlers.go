package handlers

import (
	"basicWebApp/pkg/renders"
	"net/http"
)

// Home is the homepage handler.
func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.gotmpl")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.gotmpl")
}
