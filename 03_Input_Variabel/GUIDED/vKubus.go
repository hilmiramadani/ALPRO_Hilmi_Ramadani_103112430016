package main

import (
	"fmt"
)

func main() {
	var sisi float64

	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scanln(&sisi)

	volume := sisi * sisi * sisi

	fmt.Println("Volume kubus adalah: ", sisi, volume)
}
