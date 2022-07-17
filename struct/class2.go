package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Eat() {
	fmt.Println("Person.Eat() ...")
}

func (p *Person) Walk() {
	fmt.Println("Person.Walk() ...")
}

type SuperMan struct {
	Person

	level int
}

func (s *SuperMan) Walk() {
	fmt.Println("SuperMan.Walk() ...")
}

func (s *SuperMan) Information() {
	fmt.Printf("name = %s\n", s.name)
	fmt.Printf("age = %d\n", s.age)
	fmt.Printf("level = %d\n", s.level)
}

func main() {
	superManLeo := SuperMan{
		Person{"Leo", 30},
		99,
	}

	superManLeo.Eat()
	superManLeo.Walk()
	superManLeo.Information()
}
