package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %v \n", i)
	}
	for j := 0; j <= 10; j++ {
		if j == 5 {
			for i, v := range "САШАРВО" {
				fmt.Printf("character %U '%c' starts at byte position %d\n", v, v, i)
			}
		} else {
			fmt.Printf("Nilai j = %v \n", j)
		}
	}
}
