package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("slice = %v\n", numbers)

	fmt.Printf("slice = %v\n", numbers[4:])

	fmt.Printf("slice = %v\n", numbers[:5])

	fmt.Printf("slice = %v\n", numbers[6:8])

	fmt.Println("=====")

	// numbers & s1 為同一位址
	s1 := numbers[6:8]

	fmt.Printf("slice = %v\n", s1)

	s1[0] = 100

	fmt.Printf("slice = %v\n", numbers)

	fmt.Println("=====")

	// use copy

	s2 := make([]int, 11)

	copy(s2, numbers)

	s2[4] = 200

	fmt.Printf("slice = %v\n", numbers)
	fmt.Printf("slice = %v\n", s2)
}
