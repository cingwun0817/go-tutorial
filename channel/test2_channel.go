package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main goroutine exit ...")

	c := make(chan int, 3)

	fmt.Println("len:", len(c), "cap:", cap(c))

	go func() {
		defer fmt.Println("child goroutine exit ...")

		fmt.Println("child goroutine run ...")

		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("child goroutine, i:", i, "len:", len(c), "cap:", cap(c))
		}

	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	time.Sleep(time.Second * 1)
}
