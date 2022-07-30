package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Port number of our website
const portNumber = ":8080"

// Home is the homepage handler.
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gotmpl")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gotmpl")
}

// renderTemplate renders the required template.
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

// main is the entry point.
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting an application on port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
