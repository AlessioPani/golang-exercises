package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "White",
		contactInfo: contactInfo{
			email:   "jimwhite@gmail.com",
			zipCode: 58784,
		},
	}

	jim.updateName("Jimmy") // Golang shortcut
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (ptrToPerson *person) updateName(newFirstName string) {
	(*ptrToPerson).firstName = newFirstName
}
