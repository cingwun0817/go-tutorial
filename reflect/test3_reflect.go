package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 3.14159

	numType := reflect.TypeOf(num)
	numValue := reflect.ValueOf(num)

	fmt.Println("type: ", numType)
	fmt.Println("value: ", numValue)
}
