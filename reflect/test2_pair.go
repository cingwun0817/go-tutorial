package main

import "fmt"

type Reader interface {
	ReaderBook()
}

type Writer interface {
	WriterBook()
}

type Book struct {
}

func (b *Book) ReaderBook() {
	fmt.Println("Reader Book func...")
}

func (b *Book) WriterBook() {
	fmt.Println("Writer Book func...")
}

func main() {
	b := &Book{}

	fmt.Printf("%T %v\n", b, b)

	var r Reader
	fmt.Printf("r before %T %v\n", r, r)
	r = b
	fmt.Printf("r after %T %v\n", r, r)

	r.ReaderBook()

	var w Writer
	fmt.Printf("w before %T %v\n", w, w)
	w = b
	fmt.Printf("w after %T %v\n", w, w)

	w.WriterBook()
}
