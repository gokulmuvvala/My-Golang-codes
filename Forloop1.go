package main

import "fmt"

func main() {
	var i int

	for i = 0; i < 10000; i++ {
		fmt.Println("I love Golang\n")
	}

	for j := 5; j < 15; j++ {
		fmt.Println(j)
	}

}
