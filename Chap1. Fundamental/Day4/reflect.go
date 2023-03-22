package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	x = "Hello, World!"

	// Validasi tipe data string pada variable x
	if reflect.TypeOf(x).Kind() == reflect.String {
		fmt.Println("x adalah string: ", x.(string))
	} else {
		fmt.Println("x bukan string")
	}

	x = 42

	// Validasi tipe data integer pada variable x
	if reflect.TypeOf(x).Kind() == reflect.Int {
		fmt.Println("x adalah integer: ", x.(int))
	} else {
		fmt.Println("x bukan integer")
	}
}
