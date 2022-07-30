package main

import (
	"net/http"
)

// Home is the homepage handler.
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gotmpl")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gotmpl")
}
