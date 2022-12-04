package main

import "fmt"

func main() {
	defer catch("main")
	a()
}

func a() {
	defer b()
	panic("a panic")
}

func b() {
	defer c()
	panic("b panic")
}

func c() {
	// defer catch("c")
	panic("c panic")
}

func catch(funcname string) {
	if r := recover(); r != nil {
		fmt.Println(funcname, "recover:", r)
	}
}
