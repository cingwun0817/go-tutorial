package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	sex  string
	age  int
}

func (u *User) Call() {
	fmt.Println("user Call func...")
	fmt.Printf("%v\n", u)
}

func main() {
	person := User{"Leo", "men", 18}

	person.Call()

	// get person type
	personType := reflect.TypeOf(person)
	fmt.Println("person type: ", personType.Name())

	// get person value
	personValue := reflect.ValueOf(person)
	fmt.Println("person value: ", personValue)

	// get type strings
	for i := 0; i < personType.NumField(); i++ {
		field := personType.Field(i)
		value := personValue.Field(i)

		fmt.Printf("%s %v %v\n", field.Name, field.Type, value)
	}

	// get type methods
	for i := 0; i < personType.NumMethod(); i++ {
		method := personType.Method(i)

		fmt.Printf("%s %v\n", method.Name, method.Type)
	}
}
