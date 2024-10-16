package main

import "fmt"

func main() {
	var bil1, bil2 int
	hasil := 1

	fmt.Println("masukkan bilangan pertama:")
	fmt.Scan(&bil1)
	fmt.Println("masukkan bilangan kedua:")
	fmt.Scan(&bil2)

	for i := 0; i < bil2; i++ {
		hasil *= bil1
	}
	fmt.Println(hasil)
}
