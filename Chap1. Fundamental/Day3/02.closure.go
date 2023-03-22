package main

import "fmt"

func main() {
	sayHi := func(name string) {
		fmt.Printf("Hi %s\n", name)
	}

	sayHi("Reza")

	var sum = func(a, b int) (result int) {
		result = a + b
		return
	}

	fmt.Println(sum(1, 2))
}
