package main

import "fmt"

var z *int

func escape() {
	a := 1
	z = &a
}

func main() {
	fmt.Println(z)
}
