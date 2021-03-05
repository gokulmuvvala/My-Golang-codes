package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // This loop will be the Good day is printed and this i values will be executed in a reverse manner
	}
	fmt.Println("Good day")
}
