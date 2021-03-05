// Map Example 1

package main

import "fmt"

func main() {

	// Declaration of map
	var mp map[string]int = map[string]int{ // var map name[key type]pointing data type

		"apple": 24,
		"Mango": 26,
		"melon": 6,
	}

	// printing map
	fmt.Println(mp)

	// accessing the maped elements
	// Updating the values
	mp["apple"] = 9000 // variable name [key name] = new element

	fmt.Println(mp)

	// Adding new key value

	mp["kivi"] = 20 // variable name[key name]= new value

	fmt.Println(mp)

	// deleting a key
	delete(mp, "kivi")

	fmt.Println(mp)

	// Checking the key is present or not
	val, ok := mp["Mango"] // val,ok := map name [key name]

	fmt.Println(val, ok)

	// Finding The length of map
	fmt.Println(len(mp)) //fmt.Println(len(map name))
}
