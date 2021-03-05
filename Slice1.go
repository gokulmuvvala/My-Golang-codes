// Creating a slice
package main

import "fmt"

func main() {

	// Slice Declaration
	var a []int = []int{5, 6, 7, 8, 9}
	fmt.Println(cap(a[1:3]))
	fmt.Println(len(a))
}
