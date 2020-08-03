package main

import "fmt"

func main() {
	// Declare a local variable
	var i int
	// Assign a value
	i = 42
	fmt.Println(i)
	// Reassign a value
	i = 10
	fmt.Println(i)
	// Another way to declare and initialize a variable
	var j int = 50
	fmt.Println(j)
	// Let Golang figure out the variable type
	p := 25
	fmt.Println(p)
	// Print with specific formatting
	// Print the type and value of the variable
	fmt.Printf("%v, %T", p, p)
	// Line Break
	fmt.Println("")
	q := 25.0
	fmt.Printf("%v, %T", q, q)
}