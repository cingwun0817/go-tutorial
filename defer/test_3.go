package main

import "fmt"

var g = 100

func f() int {
	defer func() {
		g = 200
	}()

	return g
}

func main() {
	i := f()

	fmt.Printf("main i = %d, g = %d\n", i, g)
}
