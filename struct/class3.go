package main

import "fmt"

type Animal interface {
	Sleep(int)
	GetColor() string
	GetType() string
}

/** Cat */
type Cat struct {
	color string
}

func (c *Cat) Sleep(time int) {
	fmt.Printf("Cat sleep %d sec\n", time)
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

/** Dog */
type Dog struct {
	color string
}

func (c *Dog) Sleep(time int) {
	fmt.Printf("Dog sleep %d sec\n", time)
}

func (c *Dog) GetColor() string {
	return c.color
}

func (c *Dog) GetType() string {
	return "Dog"
}

func showAnimalInformation(animal Animal, time int) {
	animal.Sleep(time)
	fmt.Printf("color = %s\n", animal.GetColor())
	fmt.Printf("type = %s\n", animal.GetType())
}

func main() {
	cat := Cat{"black"}
	dog := Dog{"yollow"}

	showAnimalInformation(&cat, 10)
	showAnimalInformation(&dog, 20)
}
