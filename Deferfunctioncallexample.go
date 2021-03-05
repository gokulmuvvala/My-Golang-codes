package main

import "fmt"

func main() {

	a()
}

func a() {

	fmt.Println("a Begins")
	defer b() // This is a defer function call it will be done and execute's when ever the a() function executes completly
	fmt.Println("a ends")
}

func b() {

	fmt.Println("in b")
}
