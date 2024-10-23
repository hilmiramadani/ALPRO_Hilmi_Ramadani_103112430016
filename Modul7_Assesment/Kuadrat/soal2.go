package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Masukkan bilangan n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i*i)
	}

	fmt.Println()
}
