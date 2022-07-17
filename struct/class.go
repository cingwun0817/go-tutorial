package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func (p *Person) Information() {
	fmt.Printf("name = %s\n", p.name)
	fmt.Printf("age = %d\n", p.age)
	fmt.Printf("sex = %s\n", p.sex)
}

func (p *Person) SetName(newName string) {
	p.name = newName
}

func main() {
	personLeo := Person{
		name: "Leo",
		age:  30,
		sex:  "man",
	}

	personLeo.Information()

	fmt.Println("======")

	personLeo.SetName("Leeeo")

	personLeo.Information()
}
