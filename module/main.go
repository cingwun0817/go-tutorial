package main

import (
	"fmt"
	"log"

	jsoniter "github.com/json-iterator/go"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	userA := Person{"Leo", 30}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonStr, err := json.Marshal(userA)
	if err != nil {
		log.Fatal("json marshal is failed")
	}

	fmt.Printf("%s\n", jsonStr)
}
