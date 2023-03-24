package main

import "fmt"

func main() {
	friends := [...]string{"ahyar", "cholid", "firda", "yusril", "kecek", "shaleh"}

	slice := friends[0:3]
	slice2 := friends[3:]
	var copySlice = make([]string, len(slice))

	copy(copySlice, slice)

	copySlice[0] = "anjim"
	slice[0] = "wow"

	fmt.Println("slice : ", friends)

	fmt.Println("slice : ", slice)
	fmt.Println("copySlice : ", copySlice)
	fmt.Println("slice2 : ", slice2)
	fmt.Println("slice len : ", len(slice))
	fmt.Println("slice cap : ", cap(slice))

	var _ = friends

	// konco := []string{"Ahyar", "Cholid", "Firda"}
	// var friends []string
	// friends = make([]string, 4)

	// friends[0] = "Handy"
	// friends[1] = "Reza"
	// friends[2] = "Alfanda"
	// friends[3] = ""

	// friends = append(friends[0:2], konco...)

	// name := konco[1:2] //slicing nih

	// fmt.Printf("Value of name : %v \n", name)

	// for i, friend := range friends {
	// 	fmt.Printf("key %d : value %v \n", i+1, friend)
	// }

	// nn := copy(konco, friends)

	// println(nn)

	// //cap & len
	// fmt.Printf("%v \n", konco)
	// fmt.Printf("Cap Friends : %d \n", cap(konco))
	// fmt.Printf("Len Friends : %d \n", len(konco))
}
