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

