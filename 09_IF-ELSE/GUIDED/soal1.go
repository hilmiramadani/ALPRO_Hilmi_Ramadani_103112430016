package main

import "fmt"

func main() {
	var umur int
	var kk string

	fmt.Println("masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Println("memiliki kk atau tidak:")
	fmt.Scan(&kk)

	if umur >= 17 && kk == "iya" {
		fmt.Println("bisa membuat ktp")
	} else {
		fmt.Println("belum bisa membuat ktp karena")
		if umur < 17 {
			fmt.Println("umur kurang dari 17 tahun.")
		} else if kk != "iya" && kk == "tidak" {
			fmt.Println("belum punya kk")
		}
	}
}
