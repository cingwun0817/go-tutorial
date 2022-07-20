package main

import "fmt"

var o *int

func main() {
	l := new(int)
	*l = 42

	m := &l
	n := &m
	o = **n

	fmt.Println(o)
}
