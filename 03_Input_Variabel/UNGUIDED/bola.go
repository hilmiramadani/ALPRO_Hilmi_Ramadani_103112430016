package main

import "fmt"

func main() {
	var r float32
	var volume float32
	var luas float32

	fmt.Println("masukkan jari jari:")
	fmt.Scanln(&r)

	volume = 3.14 * r * r * r / 3 * 4
	luas = 4 * 3.14 * r * r

	fmt.Println("Bola dengan jari-jari", r, "memiliki volume", volume, "dan luas kulit", luas)
}
