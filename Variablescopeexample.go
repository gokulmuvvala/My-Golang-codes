// Using global varibles (package varibles) inside the fuctions

package main

import "fmt"

var a = 9 // declaring the varible golbaly
// declaring the function definition
func demo() {
	var a = 10
	fmt.Println("From demo", a)
}

// declarinf the main function
func main() {
	demo() // calling the function
	fmt.Println("from main", a)
}
