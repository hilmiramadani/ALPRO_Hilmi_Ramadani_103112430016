package main

import "fmt"

func main() {
	var n int
	hasil := 0

	fmt.Println("masukkan n:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		hasil += i
	}
	fmt.Println(hasil)
}
