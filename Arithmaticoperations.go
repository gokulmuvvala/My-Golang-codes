package main

// Arithamatic operations

import "fmt"

func main() {
	// Declaring the varaibles here
	var a, b int

	// assigning values
	a = 5
	b = 20

	//addition
	c := a + b
	fmt.Printf("%d + %d = %d \n", a, b, c)

	// subtraction
	d := a - b
	fmt.Printf("%d - %d = %d \n", a, b, d)

	// divison
	e := float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f \n", a, b, e)

	// multiplication

	f := a * b
	fmt.Printf("%d * %d = %d \n", a, b, f)

}
