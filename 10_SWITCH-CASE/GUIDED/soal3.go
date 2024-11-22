package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, tarif int

	fmt.Println("masukan jenis kendaraan (Motor/Mobil/Truk):")
	fmt.Scanln(&kendaraan)
	fmt.Println("masukan durasi parkir (dalam jam)")
	fmt.Scanln(&durasi)

	switch {
	case kendaraan == "Motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case kendaraan == "Motor" && durasi > 2:
		tarif = 9000
	case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case kendaraan == "Mobil" && durasi > 2:
		tarif = 20000
	case kendaraan == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case kendaraan == "Truk" && durasi > 2:
		tarif = 35000
	default:
		fmt.Println("jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif parkif: Rp %d\n", tarif)
}
