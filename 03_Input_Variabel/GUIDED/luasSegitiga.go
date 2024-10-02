package main

import (
	"fmt"
)

func main() {
	var alas, tinggi float64

	fmt.Print("Masukkan panjang alas segitiga: ")
	fmt.Scanln(&alas)

	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)

	luas := 0.5 * alas * tinggi

	fmt.Println("Luas segitiga adalah:", luas)
}
