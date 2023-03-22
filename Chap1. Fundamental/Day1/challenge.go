package main

import "fmt"

func challenge() {
	i := 21
	j := true
	var k float64 = 123.456

	fmt.Printf("%d\n", i) //fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")     //fmt.Println("%")
	fmt.Printf("%t\n", j)  //    fmt.Println(j)
	fmt.Printf("%t\n", !j) // 	fmt.Println(!j)
	fmt.Printf("%c\n", 'Я')
	fmt.Printf("%d\n", 21) //fmt.Println(21)
	fmt.Printf("%o\n", 21)
	fmt.Printf("%x\n", 15)
	fmt.Printf("%X\n", 15)
	fmt.Printf("%U\n", 'Я')
	fmt.Printf("%.3f\n", k)
	fmt.Printf("%.6f\n", k)
	fmt.Printf("%E\n", k)

}
