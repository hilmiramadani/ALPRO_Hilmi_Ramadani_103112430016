package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Angka adalah Genap")
	} else {
		fmt.Println("Angka adalah Ganjil")
	}
}
