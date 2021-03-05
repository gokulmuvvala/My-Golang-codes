package main

import (
	"fmt"  // This is fmt package
	"time" // This is time package
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday() // time.now indicates current time
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
