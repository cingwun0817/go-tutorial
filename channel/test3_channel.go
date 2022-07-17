package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main goroutine exit ...")

	c := make(chan int)

	fmt.Println("len:", len(c), "cap:", cap(c))

	go func() {
		defer fmt.Println("child goroutine exit ...")

		fmt.Println("child goroutine run ...")

		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("child goroutine, i:", i, "len:", len(c), "cap:", cap(c))
		}

		close(c)
	}()

	// for {
	// 	if num, ok := <-c; ok {
	// 		fmt.Println(num)
	// 	} else {
	// 		break
	// 	}
	// }

	// use range get channel data
	for num := range c {
		fmt.Println(num)
	}

	time.Sleep(time.Second * 1)
}
