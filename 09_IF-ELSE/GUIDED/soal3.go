package main

import "fmt"

func main() {
	var (
		bilangan, angka1, angka2, angka3, angka4 int
	)

	fmt.Print("Masukkan bilangan (4 digit): ")
	fmt.Scan(&bilangan)

	if bilangan >= 1000 && bilangan <= 9999 {
		angka1 = bilangan / 1000
		angka2 = (bilangan % 1000) / 100
		angka3 = (bilangan % 100) / 10
		angka4 = bilangan % 10

		if angka1 < angka2 && angka2 < angka3 && angka3 < angka4 {
			fmt.Println("digit terurut secara membesar", bilangan)
		} else if angka1 > angka2 && angka2 > angka3 && angka3 > angka4 {
			fmt.Println("digit terurut mengecil", bilangan)
		} else {
			fmt.Println("digit tidak terurut", bilangan)
		}
	} else {
		fmt.Println("Bilangan harus terdiri dari 4 digit.")
	}

}
