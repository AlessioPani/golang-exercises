package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println("Filename:", filename)

	file, error := os.Open(filename)

	if error == nil {
		io.Copy(os.Stdout, file)
	} else {
		fmt.Println("Error:", error)
	}
}
