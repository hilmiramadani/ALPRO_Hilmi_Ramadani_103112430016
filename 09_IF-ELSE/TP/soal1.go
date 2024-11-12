package main

import (
	"fmt"
)

func main() {
	var nilai float64
	fmt.Print("Masukkan nilai mahasiswa: ")
	fmt.Scan(&nilai)

	var indeks string

	if nilai > 90 {
		indeks = "A"
	} else if nilai >= 80 && nilai <= 90 {
		indeks = "AB"
	} else if nilai >= 70 && nilai < 80 {
		indeks = "B"
	} else {
		indeks = "C"
	}

	fmt.Println("Mahasiswa mendapatkan indeks:", indeks)
}
