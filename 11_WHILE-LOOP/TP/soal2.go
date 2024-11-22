package main

import "fmt"

func main() {
	var barang string
	var harga int
	var total int

	fmt.Println("Selamat datang di aplikasi kasir sederhana!")

	for {
		fmt.Print("Masukkan nama barang (ketik 'selesai' untuk mengakhiri): ")
		fmt.Scanln(&barang)

		if barang == "selesai" {
			break
		}

		fmt.Print("Masukkan harga barang: Rp ")
		fmt.Scanln(&harga)

		total += harga

		fmt.Println("Barang", barang, "dengan harga Rp", harga, "ditambahkan.")
	}

	fmt.Println("\nTotal belanja: Rp", total)
}
