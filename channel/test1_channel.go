package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main goroutine exit ...")

	c := make(chan int)

	go func() {
		defer fmt.Println("child goroutine exit ...")

		fmt.Println("child goroutine run ...")

		for i := 0; i < 10; i++ {
			c <- i
		}

	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	time.Sleep(time.Second * 1)
}
