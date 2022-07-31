package renders

import (
	"fmt"
	"html/template"
	"net/http"
)

// tc is the template cache
var tc = make(map[string]*template.Template)

// RenderTemplateOld renders the required template loading it always from local disk.
func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gotmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

// RenderTemplate renders the required template loading it from cache, if it is there.
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		fmt.Println("creating template and reading from cache")
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// we have the template in the cache
		fmt.Println("using cached template...")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// createTemplateCache populates the template cache
func createTemplateCache(t string) error {
	templates := []string{
		"./templates/" + t,
		"./templates/base.layout.gotmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// we add template in cache
	tc[t] = tmpl
	return nil
}
