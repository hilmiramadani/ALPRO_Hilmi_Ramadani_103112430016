package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan jumlah baris segitiga bintang: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
