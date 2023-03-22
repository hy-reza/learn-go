package main

import (
	"fmt"
	"os"
)

type teman struct {
	absen           int
	nama            string
	alamat          string
	pekerjaan       string
	alasanMemilihGo string
}

var dataTeman = []teman{
	{1, "Joko", "Jakarta", "Software Engineer", "Golang mudah dipelajari dan performanya cepat"},
	{2, "Siska", "Surabaya", "Data Scientist", "Golang cocok digunakan untuk pengolahan data besar"},
	{3, "Ahmad", "Bandung", "Frontend Developer", "Golang memiliki sintaksis yang mudah dipahami"},
	{4, "Budi", "Yogyakarta", "Backend Developer", "Golang memiliki garbage collector yang efisien"},
	{5, "Dewi", "Makassar", "Mobile Developer", "Golang cocok digunakan untuk membuat aplikasi mobile"},
	{6, "Eka", "Medan", "DevOps Engineer", "Golang mudah digunakan untuk membuat alat-alat sistem"},
	{7, "Fajar", "Bali", "Database Administrator", "Golang mudah diintegrasikan dengan basis data"},
	{8, "Gina", "Semarang", "UI/UX Designer", "Golang memungkinkan pengembangan aplikasi yang scalable"},
	{9, "Hafiz", "Palembang", "Machine Learning Engineer", "Golang memiliki library machine learning yang lengkap"},
	{10, "Ivan", "Depok", "Full Stack Developer", "Golang cocok digunakan untuk membuat aplikasi web"},
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Silahkan masukkan nomor absen teman yang ingin ditampilkan.")
		return
	}

	absen := args[0]
	for _, teman := range dataTeman {
		if fmt.Sprintf("%d", teman.absen) == absen {
			fmt.Println("Nama :", teman.nama)
			fmt.Println("Alamat :", teman.alamat)
			fmt.Println("Pekerjaan :", teman.pekerjaan)
			fmt.Println("Alasan memilih Golang :", teman.alasanMemilihGo)
			return
		}
	}

	fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan.")
}
