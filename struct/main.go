package main

import "fmt"

type payload struct {
	columns []column
}

type column struct {
	source_field string
	field        string
}

func main() {
	payload := payload{
		columns: []column{
			{
				source_field: "A",
				field:        "B",
			},
			{
				source_field: "A",
				field:        "B",
			},
			{
				source_field: "A",
				field:        "B",
			},
			{
				source_field: "A",
				field:        "B",
			},
		},
	}

	fmt.Println(payload)
}
