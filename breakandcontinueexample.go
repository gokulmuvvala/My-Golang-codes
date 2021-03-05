package main

import "fmt"

func main() {

	var i int
	for i = 0; i < 5; i++ {

		if i == 3 {

			break
		}
		fmt.Println(i)
	}

	for j := 5; j < 11; j++ { // Here j is not declared as vairable so we use := for assigning the values
		if j == 7 {

			continue
		}
		fmt.Println(j)
	}
}
