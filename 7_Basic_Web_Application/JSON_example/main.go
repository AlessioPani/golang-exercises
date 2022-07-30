package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	// Write a Struct from a JSON
	myJSON := `
    [
        {
            "first_name": "Clark",
            "last_name": "Kent",
            "hair_color": "black",
            "has_dog": true
        },
        {
            "first_name": "Bruce",
            "last_name": "Wayne",
            "hair_color": "black",
            "has_dog": false
        }
    ]
    `

	unmarshalled := []Person{}

	err := json.Unmarshal([]byte(myJSON), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling a JSON", err)
	}

	fmt.Printf("umarshalled: %v", unmarshalled)

	// Write a JSON from a Struct
	mySlice := []Person{}
	firstSlice := Person{
		FirstName: "John",
		LastName:  "Wayne",
		HairColor: "red",
		HasDog:    true,
	}
	secondSlice := Person{
		FirstName: "Kentaro",
		LastName:  "Miura",
		HairColor: "black",
		HasDog:    false,
	}

	mySlice = append(mySlice, firstSlice, secondSlice)

	newJSON, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		fmt.Println("Error mashalling:", err)
	}

	fmt.Println(string(newJSON))
}
