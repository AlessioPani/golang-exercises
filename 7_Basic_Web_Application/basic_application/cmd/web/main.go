package main

import (
	"basicWebApp/pkg/config"
	"basicWebApp/pkg/handlers"
	"basicWebApp/pkg/renders"
	"fmt"
	"log"
	"net/http"
)

// Port number of our website
const portNumber = ":8080"

// main is the entry point.
func main() {
	var app config.AppConfig

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplates(&app)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Starting an application on port", portNumber[1:])
	err = serve.ListenAndServe()
	log.Fatal(err)

}
