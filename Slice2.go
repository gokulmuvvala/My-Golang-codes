// Creating a slice
package main

import "fmt"

func main() {
	var a [5]int = [5]int{5, 6, 7, 8, 9}
	// Slice Declaration
	var s []int = a[1:3]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
