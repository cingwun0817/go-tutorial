package main

import "fmt"

func main() {
	defer fmt.Println("main goroutine exit ...")

	c := make(chan int)
	quit := make(chan int)

	go func() {
		defer fmt.Println("child goroutine exit ...")

		fmt.Println("child goroutine run ...")

		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	// use select monitor more channels
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
