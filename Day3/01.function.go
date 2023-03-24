package main

import "fmt"

// func sayHello() {
// 	fmt.Printf("Hello !!")
// }
func sayHello(name string) {
	fmt.Printf("Hello, my name is %s", name)
}

func main() {
	// sayHello("Funcy \n")
	// fmt.Printf("%d \n", sum(1, 2))
	// fmt.Printf("%d \n", sum(100, 100))

	// aa, bb := freak(1, "b")
	// fmt.Printf("%d %s \n", aa, bb)

	// listFriends("Handy", "Reza", "Alfanda", "Cholid", "Firdaus", "Abada")
	// plus(sum, 2)
	fmt.Println(factorial(5))

}

//function in GO
func sum(a, b int8) (result int8) {
	result = a + b
	return
}

//multiple return value
func freak(a int, b string) (aa int, bb string) {
	aa = a + a
	bb = b + b
	return
}

func listFriends(names ...string) {
	for i, name := range names {
		fmt.Println(i, name)
	}
}

func plus(sum func(int8, int8) int8, num int8) {
	fmt.Println(sum(num, num) * num)
}

func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		println(value)
		return value * factorial(value-1)

	}
}
