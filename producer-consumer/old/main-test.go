package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var consumerCount int = 2

	queue := make(chan *data)
	wg.Add(consumerCount)

	go producer(queue)

	for i := 0; i < consumerCount; i++ {
		go consumer(i, queue)
	}

	wg.Wait()
}

type data struct {
	count int
}

func producer(ch chan<- *data) {
	defer close(ch)

	for i := 0; i < 100; i++ {
		ch <- &data{i}
	}
}

func consumer(worker int, qch <-chan *data) {
	defer wg.Done()

	for row := range qch {
		fmt.Println(worker, row.count)
	}
}
