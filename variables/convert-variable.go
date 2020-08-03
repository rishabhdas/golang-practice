package main

import "fmt"

func main() {
	// Declare an Integer type variable
	var i int = 50
	fmt.Printf("%v, %T\n", i, i)
	// Convert the above variable to Float
	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)
}