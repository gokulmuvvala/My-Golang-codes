package main

import "fmt"

func main() {

	num := 4 // declaring varible
	switch num {
	// No need of using break statemnet in ending of case in Golang
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")

	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("None")
	}

}
