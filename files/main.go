package main

import (
	"fmt"
	"io"
	"os"
)
func main() {
	fmt.Println("Welcome to files in golang")

	content := "Content file"

	file, err := os.Create("./my_file.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)

	defer file.Close()
}