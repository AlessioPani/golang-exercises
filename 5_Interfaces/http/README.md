# Interface - Example using http package



> Create an application that prints out the body of an http request.



## Body

<img src="img\resp_struct.png" alt="resp_struct" style="zoom:80%;" />

Why does Body is defined as an interface? Because if there are several data source we can reuse the same logic without code to be rewritten. Basically every data source is going to implementing the Reader interface.



### First version

```go
bs := make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
```



### Second version

<img src="img\body_interface.png" alt="body_interface" style="zoom:80%;" />



```go
io.Copy(os.Stdout, resp.Body)
```



<img src="img\iocopy.png" alt="iocopy" style="zoom:80%;" />



## Custom Writer

```go
type logWriter struct{}

func main() {
    
    ...

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "bytes.")
	return len(bs), nil
}
```

