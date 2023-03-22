package main

import (
	"fmt"
)

type Vehicle interface {
	fly()
}

type Plane struct {
	airline, model string
}

func (p Plane) fly() {
	fmt.Println(p.airline, p.model, "Fly")
}

func doSomething() interface{} {
	return "STRING"
}

func main() {
	garuda := Plane{airline: "Garuda Indonesia", model: "IDFLY19"}
	thing := doSomething()

	fmt.Println(thing.(string))

	garuda.fly()
}
