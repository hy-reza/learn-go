package main

import (
	"fmt"
	"reflect"
)

func main() {
	age := 20
	value := reflect.ValueOf(age)
	typeOf := reflect.TypeOf(age)

	fmt.Printf("value : %v %v \n", value, typeOf)

	if value.Kind() == reflect.Int {
		fmt.Println("betul")
	}
}
