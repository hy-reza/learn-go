package main

import (
	"fmt"
	"reflect"
)

func validateInput(input interface{}) error {
	inputType := reflect.TypeOf(input)

	if inputType.Kind() != reflect.String {
		return fmt.Errorf("input harus berupa string, tetapi input berupa %v", inputType.Kind())
	}

	return nil
}

func main() {
	err := validateInput("hello")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("input valid")
	}
}
