package main

import "fmt"

func main() {

	type myMap = map[string]string

	person := myMap{"name": "reza"}
	var num uint8 = 1 //sane
	var num2 byte = 1 //same

	fmt.Printf("%T, %T, %v", num, num2, person)
}
