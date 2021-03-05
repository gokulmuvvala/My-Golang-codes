package main

import "fmt"

func main() {

	var a = 4 // declareing the variables

	// increment

	fmt.Printf("a = %d \n", a)
	a = a + 1
	fmt.Printf("a + 1 = %d \n", a)
	a++
	fmt.Printf("a++ = %d \n", a)

	// Decrement
	a = a - 1
	fmt.Printf("a - 1 = %d \n", a)
	a--
	fmt.Printf("a-- = %d \n", a)

}
