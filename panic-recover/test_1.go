package main

import "fmt"

func a() {
	defer fmt.Println("defer a")
	b()
	fmt.Println("after a")
}

func b() {
	// defer fmt.Println("defer b")
	defer func() {
		fmt.Println("defer b")
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v\n", x)
		}
	}()
	c()
	fmt.Println("after b")
}

func c() {
	defer fmt.Println("defer c")
	panic("this is panic")
	fmt.Println("after c")
}

func main() {
	a()
}
