package main

import "fmt"

func main() {
	var x int // variable declaration
	x = 10    // assigning 10 to x

	fmt.Println(x)  // print a value of x
	fmt.Println(&x) // print address of x

	// declare pointer
	var num *int // num is a pointer

	val := new(int) // val varible declared
	num = new(int)
	*num = x // set value
	// & will points address  of x
	val = &x // set address

	fmt.Println("===pointer num===")
	fmt.Println(*num) // print a value of x
	fmt.Println(num)  // print address of x
	fmt.Println("===pointer val===")
	fmt.Println(*val) // print a value of x
	fmt.Println(val)  // print address of x

}
