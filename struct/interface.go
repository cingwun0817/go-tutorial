package main

import "fmt"

type Book struct {
	name string
}

func myFunc(arg interface{}) {
	fmt.Println("Call myFunc ...")
	fmt.Printf("type = %T, value = %v\n", arg, arg)

	_, ok := arg.(string)
	if ok {
		fmt.Println("Is string type")
	} else {
		fmt.Println("Is not string type")
	}
}

func main() {
	book := Book{"PHP"}

	myFunc(book)

	myFunc("AAA")

	myFunc(100)

	myFunc(3.14159)
}
