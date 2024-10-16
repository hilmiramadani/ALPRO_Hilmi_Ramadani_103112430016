package main

import "fmt"

func main() {
	var a, b, hasil int

	fmt.Println("masukkan bilangan pertama dan kedua")
	fmt.Scan(&a, &b)

	for i := 1; i <= b; i += 1 {
		hasil = hasil + a
	}
	fmt.Println(hasil)
}
