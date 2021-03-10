// Employes data collection using Golang

package main

import "fmt"

// Structure declaration
type Employee struct {
	Employeename string
	Age          int
	Employeeid   int
	Cabinnumber  int
}

func main() {
	// Employee 1 data
	var E1 Employee = Employee{"Gokul", 20, 183, 24}
	fmt.Println(E1)

	// Employee 2 data
	var E2 Employee = Employee{"Kumar", 21, 168, 97}
	fmt.Println(E2)

}
