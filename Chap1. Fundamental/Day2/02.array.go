package main

import "fmt"

func main() {
	var cities [5]string
	cities[0] = "Malang"
	cities[1] = "Surabaya"
	cities[2] = "Jember"

	friends := [3]string{"ahyar", "cholid", "firda"}

	for i, friend := range friends {
		fmt.Println(i+1, ":", friend)
	}
	for i, city := range cities {
		fmt.Println(i+1, ":", city)
	}

	fmt.Println("panjang array kota :", len(cities))
	fmt.Println("panjang array teman :", len(friends))
	fmt.Println("kapasitas array kota :", cap(cities))
	fmt.Println("kapasitas array teman :", cap(friends))

	// //ARRAY
	// // var name = [5]string{"Handy", "Reza", "Alfanda", "Cholid", "Firdaus"}

	// // name[4] = "Abada"

	// // fmt.Printf("My Friend %v", name)

	// // for i, friend := range name {
	// // 	fmt.Printf("%d : %s \n", i, friend)
	// // }

	// var name = [2][2]string{{"Handy", "Reza"}, {"Cholid", "Firdaus"}}

	// name[1][1] = "Abada"

	// fmt.Printf("My Friend %v \n", name)

	// for _, friend := range name {
	// 	for i, f := range friend {
	// 		fmt.Printf("%d : %s \n", i, f)
	// 	}
	// }
}
