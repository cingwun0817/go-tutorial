package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func readPageSize(url string, channel chan Page) {
	fmt.Println("Get " + url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{Url: url, Size: len(body)}
}

func greeting(myChannel chan string) {
	myChannel <- "Hi"
}

func abc(channel chan string) {
	channel <- "A"
	channel <- "B"
	channel <- "C"
}

func def(channel chan string) {
	channel <- "D"
	channel <- "E"
	channel <- "F"
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second * 1)
	}

	fmt.Println(name, "wakes up!!")
}

func send(channel chan string) {
	reportNap("send goroutine", 2)

	fmt.Println("** send value: A **")
	channel <- "A"

	fmt.Println("** send value: B **")
	channel <- "B"
}

type Page struct {
	Url  string
	Size int
}

func main() {
	// channel := make(chan string)
	// go send(channel)

	// reportNap("main goroutine", 5)

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// channelA := make(chan string)
	// channelB := make(chan string)

	// go abc(channelA)
	// go def(channelB)

	// fmt.Println(<-channelA)
	// fmt.Println(<-channelA)
	// fmt.Println(<-channelA)
	// fmt.Println(<-channelB)
	// fmt.Println(<-channelB)
	// fmt.Println(<-channelB)

	// go a()
	// go b()

	// pages := make(chan Page)
	// urls := []string{"https://www.google.com/", "https://www.atelli.ai/en/", "https://www.adgeek.com/en/"}

	// for _, url := range urls {
	// 	go readPageSize(url, pages)
	// }

	// for i := 0; i < len(urls); i++ {
	// 	page := <-pages

	// 	fmt.Println(page.Url, page.Size)
	// }

	// myChannel := make(chan string)
	// go greeting(myChannel)

	// output := <-myChannel
	// fmt.Println(output)

	// time.Sleep(time.Second)

	// fmt.Println("end")

	link := make(chan string)
	done := make(chan bool)
	go producer(link)

	for i := 1; i <= 10; i++ {
		go consumer(i, link, done)
	}

	fmt.Println("end", <-done)
}

var messages = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
	"- Elliot Alderson",
	"Mr. Robot",
}

func producer(link chan<- string) {
	for _, m := range messages {
		link <- m
	}
	close(link)
}

func consumer(worker int, link <-chan string, done chan<- bool) {
	for b := range link {
		fmt.Println(worker, b)

		time.Sleep(3 * time.Second)
	}
	done <- true
}
