package main

import "fmt"

// Define single package variable
var i float32 = 50.5

// Define a block of variables
var (
	name string = "Rishabh"
	age int = 10
)

// := syntax is not supported for package level variables

func main() {
	fmt.Println(i)
	fmt.Println(name)
	fmt.Println(age)
}