package main

import "fmt"

func main() {

	//LOOPINGS

	// 	for i := 0; i <= 10; i++ {
	// 		fmt.Printf("Angka %d \n", i)
	// 	}

	// 	j := 0
	// 	for j < 3 { //pengganti While
	// 		fmt.Print(j)
	// 		j++
	// 	}
	// 	for { //pengganti While
	// 		j++
	// 		if j%2 == 1 {
	// 			continue
	// 		}
	// 		if j >= 10 {
	// 			break
	// 		}
	// 		fmt.Print(j)
	// 	}

	// outerLoop:
	// 	for x := 0; x < 100; x++ {
	// 		// innerLoop:
	// 		fmt.Printf("x: %d \n", x)
	// 		for y := 0; y < 2; y++ {
	// 			if x == 10 {
	// 				break outerLoop
	// 			}
	// 		}
	// 	}

	// cars := []string{"ferari", "pors", "nissan", "bmw", "lamborghini"}

	// fmt.Println("----------------")
	// for i := 0; i < len(cars); i++ {
	// 	fmt.Println(i+1, cars[i])
	// }

	// fmt.Println("----------------")

	// for i, car := range cars {
	// 	fmt.Println(i+1, car)
	// }

	for i := 0; i < 50; i++ {
		if i%2 != 0 {
			continue
		} else if i > 20 {
			break
		}
		fmt.Println(i)
	}

}
