package main

import (
	"fmt"
)

func main() {
	var nilaiTukar float64 = 15000.0

	var uangIDR int

	fmt.Print("Masukkan jumlah uang dalam IDR: ")
	fmt.Scanln(&uangIDR)

	uangUSD := float64(uangIDR) / nilaiTukar

	fmt.Println("Jumlah uang dalam USD adalah:", uangUSD)
}
