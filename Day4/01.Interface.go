package main

import "fmt"

type Animal interface {
	Speak()
}

type Cat struct{}

func (c *Cat) Speak() {
	fmt.Println("Meow!")
}

type Dog struct{}

func (d *Dog) Speak() {
	fmt.Println("Woof!")
}

func main() {
	var animal Animal

	animal = &Cat{}
	animal.Speak() // Output: Meow!

	animal = &Dog{}
	animal.Speak() // Output: Woof!
}
