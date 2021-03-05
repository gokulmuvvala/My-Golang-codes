// Creating a slice
package main

import "fmt"

func main() {

	// Array Declaration
	var a []int = []int{5, 6, 7, 8, 9}

	b := append(a, 20, 36)
	fmt.Println(b)

}
