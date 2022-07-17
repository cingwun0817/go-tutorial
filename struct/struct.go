package main

import "fmt"

type Book struct {
	title       string
	description string
}

func changeTitle(book Book) {
	book.title = "Name changed"
}

func changeDesc(book *Book) {
	book.description = "Desc changed"
}

func main() {
	var book1 Book

	book1.title = "Name"
	book1.description = "Desc"

	fmt.Printf("%v\n", book1)

	// change title
	changeTitle(book1)

	fmt.Printf("%v\n", book1)

	// change description, use point
	changeDesc(&book1)

	fmt.Printf("%v\n", book1)
}
