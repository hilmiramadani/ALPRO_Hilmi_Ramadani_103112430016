package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan angka N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i*i)
	}
	fmt.Println()
}
