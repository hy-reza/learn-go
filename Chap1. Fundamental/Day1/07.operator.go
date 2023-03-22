package main

import "fmt"

// +, -, *, /, % <-- Aritmatic
// ==, !=, >, <, >=+, <= <-- Logic
// &&, !. || <-- bolean
func main() {
	a := 10
	b := 10
	result := a + b

	result += a + b/2

	result++
	fmt.Println(result)
	fmt.Println(result%2 == 0)
	fmt.Println(result%2 != 0)
	fmt.Println(result%2 > 0)
	fmt.Println(result%2 < 0)

	fmt.Println(true && !true)
	fmt.Println(true && !false)

	const nilai = 90
	const kehadiran = 100

	const lulusUjian = nilai >= 75
	const lulusKehadiran = kehadiran >= 85

	const lulus = lulusKehadiran && lulusUjian

	fmt.Println("lulus : ", lulus)

}
