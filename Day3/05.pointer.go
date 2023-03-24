package main

import "fmt"

func addArr(arr *[]string, str string) {
	*arr = append(*arr, str)
}

func main() {

	arr := []string{"satu", "dua", "tiga"}

	fmt.Println(arr)
	addArr(&arr, "empat")
	fmt.Println(arr)
	// num1 := 4
	// num2 := &num1 // & => alamat

	// println(num1)
	// println(*num2) // * => value dari alaPrintln()
	// println(&num1)
	// println(num2)

	// *num2 = num1 * 5

	// println(num1)

	// changeValue(&num1)
	// println(num1)
}

// use case
// func changeValue(num *int) {
// 	*num = 99
// }
