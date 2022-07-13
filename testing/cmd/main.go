package main

import (
	"fmt"

	compute "test/pkg"
)

func main() {
	valOne := compute.Plus(1, 3)
	varTwo := compute.Times(5, 8)

	fmt.Println(valOne, varTwo)
}
