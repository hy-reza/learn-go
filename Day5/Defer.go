package main

import "fmt"

func divide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	return a / b, nil
}

func main() {

}
