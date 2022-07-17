package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("main exit ...")

	fmt.Println("main goroutine ...")

	go func() {
		defer fmt.Println("child exit ...")

		fmt.Println("child goroutine ...")

		func() {
			defer fmt.Println("child content func exit ...")

			runtime.Goexit() // 退出當前 goroutine

			return
		}()

		j := 0
		for {
			j++
			fmt.Println("child j:", j)

			time.Sleep(time.Second * 1)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main i:", i)

		time.Sleep(time.Second * 1)
	}
}
