package main

import "fmt"

// Function declaration
// Creating our own add function
func calc(x, y int) (int, int) { // fun function name(varibles)return values datatypes
	var out1 = x + y
	var out2 = x - y
	return out1, out2 // return what you want to return (It can also return 2 values glang special feauture)

}

func main() {

	a := 24
	b := 69
	result1, result2 := calc(a, b) // calling function
	fmt.Print(result1, result2)
}
