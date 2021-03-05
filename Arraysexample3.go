package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("define arrays")
	var numbers [5]int   // 1D array with 5 length
	var cities [5]string // 1D array with 5 length
	var matrix [3][3]int // 2D array with 3 arrays and 3 elements

	// Inseet data

	fmt.Println(">>>>>>> Inset array data ")
	for i := 0; i < 5; i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("City %d", i)
	}

	//insert matrix data

	fmt.Println(">>>>>>> Inset matrix data ")
	for i := 0; i < 3; i++ { // this for loop for 3 arrays
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}
	// display data
	fmt.Println(">>>>>>> display array data ")
	for j := 0; j < 5; j++ {
		fmt.Println(numbers[j])
		fmt.Println(cities[j])
	}

	// display matrix data
	fmt.Println(">>>>>>> display matrix data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matrix[i][j])
		}
	}

}
