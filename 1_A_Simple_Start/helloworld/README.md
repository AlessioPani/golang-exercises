# First application



## How do we run the code in our project

To run the code simply get (using a terminal) to the folder where the code is located and write the following instruction:

```powershell
go run filename.go
```

This instruction will compile and execute your code. If you want to only compile and build an executable use instead this instruction:

```powershell
go build filename.go
```

In order to create a package, use this instruction:

```powershell
go mod init github.com/YourUserName/module name
```



## What does 'package main' mean

A package can be seen as a project or a  workspace and of course can be made of several .go files. 
With *package main* we are saying that those go files are related to a single package called “main”.

So we could use a random name instead of *main*? Well, it depends.
In Go there are two different type of package: an **executable** that generates an .exe to be run (using the previous *build* command) and a **reusable** code, used like a library or a dependency. 

* If you want to create an executable package, you have to use the **keyword main** to tell the compiler that you actually want an .exe and it must be a function called **main** in it;
* If you want to create a reusable code, you can choose a random name to call the package. The compiler will generate nothing.



## What does 'import "fmt"' mean

The **import** statement give us access to all of the code and functionality in the “fmt” package, which is a standard library ([here](https://pkg.go.dev/std) you can find the official page).



## How is the main.go file organized

| **package main**      | Package declaration                         |
| :-------------------- | :------------------------------------------ |
| **import “fmt”**      | **Import other packages we need**           |
| **func main() { … }** | **Declare functions, tell Go to do things** |

