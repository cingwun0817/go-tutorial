package main

import (
	"fmt"
	"sync"
	"time"
)

type Tag struct {
	TagId   string
	TagName string
}

func main() {
	smap := sync.Map{}

	smap.Store("123", Tag{TagId: "123", TagName: "AAA"})

	go func() {
		i := 0
		for i < 1000 {
			smap.Store("456", Tag{TagId: "456", TagName: "BBB"})
			i++
		}
	}()

	go func() { //19 行
		i := 0
		for i < 1000 {
			smap.Store("456", Tag{TagId: "456", TagName: "CCC"})
			i++
		}
	}()

	time.Sleep(time.Second * 1)

	val, ok := smap.Load("456")
	tag := val.(Tag)
	fmt.Println(tag.TagId, tag.TagName, ok)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)

		return true
	})
}
