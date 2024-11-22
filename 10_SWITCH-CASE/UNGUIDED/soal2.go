package main

import (
	"fmt"
	"math"
)

func main() {
	var jenisKendaraan string
	var durasiParkir float64
	var tarifPerJam, totalBiaya int

	fmt.Println("Masukkan jenis kendaraan (motor/mobil/truk):")
	fmt.Scanln(&jenisKendaraan)

	fmt.Println("Masukkan durasi parkir dalam jam:")
	fmt.Scanln(&durasiParkir)

	durasiParkir = math.Ceil(durasiParkir)

	switch jenisKendaraan {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid!")
		return
	}

	totalBiaya = int(durasiParkir) * tarifPerJam

	fmt.Printf("Jenis Kendaraan: %s\n", jenisKendaraan)
	fmt.Printf("Durasi Parkir  : %.0f jam\n", durasiParkir)
	fmt.Printf("Total Biaya    : Rp %d\n", totalBiaya)
}
