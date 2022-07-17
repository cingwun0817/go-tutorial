package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `info:"name" doc:"my name"`
	age  int    `info:"age"doc:"my age"`
}

func main() {
	var userA Person

	userAType := reflect.TypeOf(userA)

	for i := 0; i < userAType.NumField(); i++ {
		infoTag := userAType.Field(i).Tag.Get("info")
		docTag := userAType.Field(i).Tag.Get("doc")

		fmt.Println("info:", infoTag, ", doc:", docTag)
	}
}
