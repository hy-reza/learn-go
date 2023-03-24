package main

import (
	"fmt"
	"strings"
)

func main() {
	// inisialisasi teks
	text := "Elit mollit non enim anim occaecat elit proident sunt excepteur id nostrud aliqua Lorem sint. Aute aliqua incididunt in excepteur ad magna tempor quis. Ex elit non laboris laboris officia sit. Dolor deserunt sit adipisicing nulla eiusmod."
	// memecah teks menjadi karakter-karakter
	words := strings.Split(text, "")
	// membuat map untuk menyimpan jumlah kemunculan karakter
	charCount := make(map[string]int)

	// melakukan perulangan pada setiap kata dalam teks
	for _, word := range words {
		// melakukan perulangan pada setiap karakter dalam kata
		for _, char := range word {
			// menambahkan karakter ke dalam map charCount dan menginkremen jumlah kemunculan
			charCount[string(char)]++
			// mencetak karakter ke konsol
			fmt.Println(string(char))
		}
	}

	// mencetak map charCount yang berisi jumlah kemunculan setiap karakter dalam teks
	fmt.Println(charCount)
}
