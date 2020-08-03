package main

import "fmt"

// Define a oackage level variable i
var i int = 50

func main() {
	// Print the global variable
	fmt.Println(i)
	// Define i in the local scope
	var i int = 25
	fmt.Println(i)
}