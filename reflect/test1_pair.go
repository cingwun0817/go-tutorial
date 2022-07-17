package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// tty: pair<type: *os.File, value: "./test.log">
	tty, err := os.OpenFile("./test.log", os.O_RDWR, 0)

	if err != nil {
		log.Fatal(err)
		return
	}

	// r: pair<type: , value: >
	var r io.Reader
	// r: pair<type: *os.File, value: "./test.log">
	r = tty

	// w: pair<type: , value: >
	var w io.Writer
	// w: pair<type: *os.File, value: "./test.log">
	w = r.(io.Writer)

	fmt.Printf("%T\n", w)

	w.Write([]byte("HELLO"))
}
