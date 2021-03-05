package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 12.5

	var result = math.Sqrt(num)

	//fmt.Println(result)
	fmt.Printf("The output %f Thank you\n", result)  //  %f default width, default precision (After . As much as languge supports)
	fmt.Printf("The output %2f Thank you\n", result) // %.2f   default width, precision 2 (After . 2 decimals )
	fmt.Printf("The output %4f Thank you\n", result) // %.2f   default width, precision 4 (After . 4 decimals  )
	fmt.Printf("The output %g Thank you", result)    // %g	 for large exponents
}
