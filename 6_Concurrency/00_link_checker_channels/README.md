# Concurrency and Go Channels



## Link status checker - Serial mode

<img src="img\status_checker1.png" alt="status_checker1" style="zoom:80%;" />



<img src="img\status_checker2.png" alt="status_checker2" style="zoom:80%;" />



<img src="img\status_checker3.png" alt="status_checker3" style="zoom:80%;" />



```go
package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.com",
		"https://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(l string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "may be down.")
		return
	}
	fmt.Println(l, "is up!")
}
```



## Link status checker - Parallel mode

<img src="img\status_checker_parallel.png" alt="status_checker_parallel" style="zoom:80%;" />

In order to accomplish this task, we can rely on **go routines** and **channels**.

When we create a Go program, we’re actually using a single **go routine** that runs on a single core. 



> What about run different tasks (blocking call calls ) on separated go routines (so on multiple cores/threads)? 

```go
_, err := http.Get(l) // blocking call!
```



### Syntax of Go routines and channels

```go
go checkLink(link) // Create a new go routine
```

<img src="img\threads.png" alt="threads" style="zoom:80%;" />



Unfortunately this simple modification is not enough to enable parallelism in our application. 

<img src="img\early_exit.png" alt="early_exit" style="zoom:80%;" />

In fact, the main routine is not aware if all the child go routine are done or not, so when the main code it’s done, the program exit entirely. 

We must use another feature to communicate to the main routine the status of a child.
We must use **Go channels**, which are typed struct. They actually send a message of a specified type.

<img src="img\channels.png" alt="channels" style="zoom:80%;" />



<img src="img\typed_channels.png" alt="typed_channels" style="zoom:80%;" />



<img src="img\sending_data_channel.png" alt="sending_data_channel" style="zoom:80%;" />



```go
c := make(chan string)
func checkLink(l string, c chan string) { ... } // Now checkLink has access to our channel
```



#### Receiving messages

```go
func main() {

	[...]
	
	fmt.Println(<-c)
}
```

<img src="img\timeline.png" alt="timeline" style="zoom:80%;" />



```go
func main() {

	[...]
	
	fmt.Println(<-c)
     fmt.Println(<-c)
}
```



<img src="img\multiple_receives.png" alt="multiple_receives" style="zoom:80%;" />



How can we read all the channels and only then exit the program entirely? Using for example a for loop. 

```go
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // reading a channel is a blocking task, so the for loop will not be executed
          			  // instantly, in fact the single iteration will wait a message in the channel
	}
```



#### Repeating routines

```go
func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.com",
		"https://amazon.com",
	}

	c := make(chan string)

	// First iteration
	for _, link := range links {
		go checkLink(link, c)
	}

	// Continuous iteration
	for {
		go checkLink(<-c, c)
	}

}

// Check the status of a link
func checkLink(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "may be down.")
		c <- l
		return
	}
	fmt.Println(l, "is up!")
	c <- l
}
```



In order to clarify what’s going on we can use an alternative and clearer for loop.

```go
// Continuous iteration
for l := range c {
	go checkLink(l, c)
}
```



#### Sleeping a Routine and Function Literals

This program actually works quite well, but the pinging of the websites is still very quick. How can add a pause between calls?
We can use the **time** standard package and one of its functions, **Sleep**.

```go
time.Sleep(time.Second) // Wait for 1 sec
```

We can’t use the Sleep fuction inside the main function in this case, because would stops the main routine. 

We can use a function literal (like Lambdas in Python).

```go
for l := range c {
	go func(link string) {
		time.Sleep(5 * time.Second) // Wait for 5 sec
		checkLink(link, c)
	}(l)
}
```

> This function literal does have the link as an input to have a reference to the new link in memory instead of a copy of the same link in memory.
>
> **Results without passing the link:**
> https://stackoverflow.com is up!
> https://google.com is up!
> https://facebook.com is up!
> https://golang.com is up!
> https://amazon.com is up!
> https://amazon.com is up!
> https://amazon.com is up!
> https://amazon.com is up!
>
> **Results with passing the link:**
>
> https://stackoverflow.com is up!
> https://google.com is up!
> https://facebook.com is up!
> https://golang.com is up!
> https://amazon.com is up!
> https://stackoverflow.com is up!
> https://google.com is up!
> https://facebook.com is up!
> https://golang.com is up!
> https://amazon.com is up!



### Concurrency vs Parallelism

<img src="img\concurrency.png" alt="concurrency" style="zoom:80%;" />



<img src="img\parallelism.png" alt="parallelism" style="zoom:80%;" />
