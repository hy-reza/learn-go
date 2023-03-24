package main

import "fmt"

//VARIABLE
//var namavar typedata = value
func main() {
	var name string
	name = "Handy Reza Alfanda"
	fmt.Println(name)
	name = "Handy Reza"
	fmt.Println(name)

	var friendName = "Cholid"
	fmt.Println(friendName)
	var friendAge uint8 = 32
	fmt.Println(friendAge)

	country := "Indonesia"
	fmt.Println(country)

	//multiple variable declaration
	// var (
	// 	food = "Meetball",
	// 	drink = "Tea"
	// )
	var food, drink string = "Chiken Briani", "Boba"
	fmt.Println(food, drink)

	//DATA TYPE CONVERTION
	var age int8 = 22
	var age64 int64 = int64(age)
	var h = name[0]
	var h8 = string(h)

	_, _, _ = age64, h, h8
}
