# Interfaces

The shuffle method does not have actually a specific reference to a deck type. If we want to use the same logic with a different receiver we have to rewrite a function.

<img src="img\shuffle.png" alt="shuffle" style="zoom:80%;" />

Using interfaces we can reuse a logic that has been already written.

## Example - Bots

<img src="img\bots.png" alt="bots" style="zoom:80%;" />



```go
type bot interface {
	getGreeting() string
}

...

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

...
```

With an **interface**, we’re actually defining a method/function set with their return type for another type must have to be considered of type bot.
An interface in Golang is **implicit**, so we don’t manually have to say that our custom type satisfies some interface.

