# Structs

A struct (short for *structure*) is a collection of properties that are related together. 

The previous project (Cards) would have been more efficient using struct, as the following diagram shows:

![struct_on_cards](img\struct_on_cards.png)



## A simple project

To get used to structs in Go we can create a new simple project.



![people](img\people.png)



### Declare a struct

```go
type person struct {
	firstName string
	lastName  string
}
```



### Create a new instance of a struct

There are a lot of ways to create a new instance of a struct:

```go
alex := person{"Alex", "Anderson"}
```

```go
alex := person{firstName: "Alex", lastName: "Anderson"}
```

```go
var anotherAlex person
```

In the last case, both firstName and lastName are empy string as a their zero values.



#### Zero values

```go
fmt.Printf("%+v", anotherAlex) //Prints fields and values in an explicit way

// Results:
{firstName: lastName:}
```

![zero_values](img\zero_values.png)



### Update a struct value

```go
alex := person{firstName: "Alex", lastName: "Anderson"}
fmt.Println(alex) // {Alex Anderson}

alex.firstName = "John"
alex.lastName = "Depp"
fmt.Println(alex) //{John Depp}
```



### Embedding structs

![embedding_structs](img\embedding_structs.png)

```go
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}
```

```go
jim := person{
	firstName: "Jim",
	lastName:  "White",
	contact: contactInfo{
		email:   "jimwhite@gmail.com",
		zipCode: 58784,
	},
}
```

```go
// Results:
{firstName:Jim lastName:White contact:{email:jimwhite@gmail.com zipCode:58784}}
```



### Structs with receiver functions

```go
type person struct {
	firstName string
	lastName  string
	contactInfo //Only the type name; instead of contact contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
```

Go passes the parameters by value, so in order to actually update a name using the updateName method we MUST use a pointer (*, which give me the value the memory address is pointing at) and its reference in the memory (&, which give the memory address of the value this variable is pointing at):

```go
jimPointer := &jim //jimPointer => memory address of the variable jim (reference to the struct)
jimPointer.updateName("Jimmy")

func (ptrToPerson *person) updateName(newFirstName string) {
	(*ptrToPerson).firstName = newFirstName // changes firstName directly at the memory address 
     							     // of the jim variable
}
```

<img src="img\pointers.png" alt="pointers" style="zoom: 80%;" />



This code can be shortened using Go shortcuts; in fact Go allows us to call a method with a pointer as a receiver using either a pointer or the type. Here’s how:

```go
jim.updateName("Jimmy")
jim.print()

func (ptrToPerson *person) updateName(newFirstName string) {
	(*ptrToPerson).firstName = newFirstName
}
```



#### Value vs Reference types

![reference_value_types](img\reference_value_types.png)

* **Value types**: Go copy their values, we **must** use a pointer if we want to change their values using a function; we need their address in memory to actually make a change.

* **Reference types**: while creating one of those types (for example a slice - see the image below), Golang creates in the memory a structure which contains the pointer to head of a slice and other informations. When we call the updateSlice function, we’re copying not its values but the initial structure. Basically we have another structure in memory that points to the same memory address.

  <img src="img\slice_update_memory.png" alt="slice_update_memory" style="zoom:80%;" />
