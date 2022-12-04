package main

import "fmt"

func executePanic() {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			fmt.Println(errMsg)
		}

		fmt.Println("This is recovery function")
	}()

	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	executePanic()

	fmt.Println("Main block is executed completely")
}
