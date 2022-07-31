package renders

import (
	"basicWebApp/pkg/config"
	"basicWebApp/pkg/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// tc is the template cache
var tc = make(map[string]*template.Template)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds default data to templates
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template using the automatic cache function
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the AppConfig
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(ok)
	}

	td = AddDefaultData(td)

	err := t.Execute(w, td)
	if err != nil {
		log.Fatal(err)
	}
}

// createTemplateCache populates the template cache - "automatic" way
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.gotmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.gotmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.gotmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// get all of the files named *.layout.gotmpl from ./templates
		matches, err := filepath.Glob("./templates/*.layout.gotmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gotmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// RenderTemplate renders the required template loading it from cache, if it is there.
// "Manual" cache.
func RenderTemplateManual(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		fmt.Println("creating template and reading from cache")
		err = createTemplateCacheManual(t)
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

// createTemplateCache populates the template cache - "manual" way
func createTemplateCacheManual(t string) error {
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

// RenderTemplateOld renders the required template loading it always from local disk.
func RenderTemplateNonOptimized(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gotmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
