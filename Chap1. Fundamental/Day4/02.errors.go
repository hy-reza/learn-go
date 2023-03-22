package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("Input tidak boleh kosong")
	}

	return a + b, nil
}

func main() {
	fmt.Println(sum(-1, -2))
}
