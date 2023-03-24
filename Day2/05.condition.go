package main

import "fmt"

func main() { //CONDITION

	// age := 17
	if age := 14; age >= 17 {
		fmt.Println("Allowed!  ")
	} else if age < 17 && age >= 15 {
		fmt.Println("to soon!  ")
	} else {
		fmt.Println("Not Allowed!")
	}

	switch point := 9; point {
	case 10, 9, 8:
		fmt.Println("A")
		fallthrough //jalankan next case
	case 7:
		fmt.Println("B")
		fallthrough //jalankan next case
	default:
		fmt.Println("C")
	}

	absen := 90
	score := 90

	switch {
	case absen >= 90 && score >= 75:
		fmt.Println("Lulus")
	default:
		fmt.Println("Tidak Lulus")
	}
}
