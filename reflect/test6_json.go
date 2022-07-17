package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"Bed Body", 1991, 200, []string{"Leo", "Eric"}}

	// struct to json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("jsonStr: %s\n", jsonStr)

	// json to struct
	myMovie := Movie{}

	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", myMovie)
}
