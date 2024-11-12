package main

import (
	"fmt"
)

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Print("Masukkan umur Anda: ")
	fmt.Scan(&umur)

	fmt.Print("Masukkan kewarganegaraan Anda: ")
	fmt.Scan(&kewarganegaraan)

	if umur >= 17 && (kewarganegaraan == "WNI" || kewarganegaraan == "wni") {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		fmt.Print("Anda tidak bisa mengikuti pemilu karena ")
		if umur < 17 {
			fmt.Println("umur Anda kurang dari 17 tahun.")
		} else if kewarganegaraan != "WNI" && kewarganegaraan != "wni" {
			fmt.Println("Anda bukan WNI.")
		}
	}
}
