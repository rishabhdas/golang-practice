package main

import "fmt"

func main() {
	var i int = 25
	// Define a variable which is not used in the program
	var j int = 50
	fmt.Println(i)
	// This program should return an error since j is not used
	// Comment out line 8 to remove the error
	// This is to ensure no unused variables exist in the code
}