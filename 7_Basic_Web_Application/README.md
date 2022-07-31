# Create Web Applications

## Handling JSON

### From Struct to JSON

```go
mySlice = append(mySlice, firstSlice, secondSlice, ...) //Created slice of struct

newJSON, err := json.MarshalIndent(mySlice, "", "    ") //From Struct to JSON

//Print JSON values
fmt.Println(string(newJSON))
```



### From JSON to Struct

```go
myJSON := `[ ... ]`// Received JSON
unmarshalled := []StructType{} // Slice of the struct type you're waiting for

err := json.Unmarshal([]byte(myJSON), &unmarshalled) //From JSON to Struct

//Print struct values
fmt.Printf("umarshalled: %v", unmarshalled)
```



## Start a Web Server

```go
http.ListenAndServe(":8080", nil) // :8080 is a port
```



### Handle an http request

```go
http.HandleFunc("/", Home)

// Home is the homepage handler.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my homepage!")
}
```

### Serve an HTML template

```go
// renderTemplate renders the required template.
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

// Home is the homepage handler.
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gotmpl")
}
```



### Use Template Layout 

To reuse a barebone html template we can define a base layout with some blocks (like content, css, js, ...)

```html
// TEMPLATE
{{define "base"}}

... HTML content ...
       
	{{block "css" .}}{{end}}
    {{block "content" .}}{{end}}
    {{block "js" .}}{{end}}

{{end}}

```

and then import that in other pages

```html
// HTML PAGE
{{template "base" .}}

{{define "content"}}
<div class="container">
    <div class="row">
        <div class="col">
            <h1>This is the homepage.</h1>
            <p>This is a paragraph.</p>
        </div>
    </div>
</div>
{{end}}
```

To be read by the backend, the renderer must be updated for example in this way:

```go
parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gotmpl")
```

### Rendering optimization

There are 2 two ways to optimize the template rendering task, instead of always load the page from disk on every request, I could have some kind of data structure as cache. This cache can be filled in a manual or an automatic way.

#### Manual 

Instead of reading from the disk every single time I could have some kind of data structure. I can store a parsed template into and then I'll check to see if the template exists in that data structure: if it does, I'll use it - if it doesn't, I'll read it from disk, parse it and then store the resulting template in that data structure.

The best data structure to use in this case is a map.

```go
// tc is the template cache
var tc = make(map[string]*template.Template)

// RenderTemplate renders the required template loading it from cache, if it is there.
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

// createTemplateCacheManual populates the template cache
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
```

#### Automatic

Similar to the manual way, but the cache is entirely built during the first rendering request.

```go
// tc is the template cache
var tc = make(map[string]*template.Template)

// RenderTemplate renders a template using the automatic cache function
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// createTemplateCache populates the template cache - "automatic" way
func createTemplateCache() (map[string]*template.Template, error) {
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
```

### Passing data to templates

### Routing

#### Pat

```go
func routes(a *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

// ...

serve := &http.Server{
	Addr:    portNumber,
	Handler: routes(&app),
}

err = serve.ListenAndServe()
```



#### Chi

Chi does have a middleware built-in.

```go
func routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

    mux.Use(middleware.Recoverer) //for example using the Recovery middleware from chi pkg

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

// ...

serve := &http.Server{
	Addr:    portNumber,
	Handler: routes(&app),
}

err = serve.ListenAndServe()
```

##### Custom middleware

