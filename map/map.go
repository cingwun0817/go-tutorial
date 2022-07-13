package main

import "fmt"

func changeValue(myMap map[string]string) {
	myMap["two"] = "JAVA"
}

func main() {
	myMap := make(map[int]string)

	// add
	myMap[0] = "A"
	myMap[3] = "A"

	fmt.Println(myMap)

	myMap2 := map[string]string{
		"one":   "PHP",
		"two":   "Golang",
		"three": "Python",
	}

	fmt.Println(myMap2)

	// update
	myMap2["three"] = "rust"

	fmt.Println(myMap2)

	// delete
	delete(myMap2, "one")

	fmt.Println(myMap2)

	fmt.Println("======")

	changeValue(myMap2)

	fmt.Println(myMap2)
}
