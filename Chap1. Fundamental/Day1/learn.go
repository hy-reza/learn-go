package main

import (
	"fmt"
)

func main() {
	//Naming Conventions Variable
	var camelCase = "Naming Conventions"
	//Variable with data type
	var name string = "Handy Reza Alfanda"
	var age float64 = 22
	//short declaration
	num := 30
	happy := true
	// Multiple Declaration Variable
	// const one, two, three int = 1, 2, 3
	one, two, three := 1, 2, "tiga"

	var _, _, _, _, _, _, _ = num, name, age, camelCase, one, two, three

	println(num, happy)
	fmt.Println("Hello World!")
	fmt.Println(`Hello World! 
	Commodo ex sunt eiusmod quis proident ut duis commodo amet labore ullamco. 
	Officia officia ad ex voluptate labore elit commodo aute qui tempor. Ipsum laboris elit est aliquip consectetur nostrud veniam laborum ea est sint eu labore. Qui non magna ea sint aliqua esse amet sunt enim id ex pariatur excepteur est.`)
	//before
	fmt.Printf("age : %.3f and happy : %t \n", age, happy) //float
	//after
	age++
	fmt.Printf("age : %.3f and happy : %t", age, !happy) //float

}
