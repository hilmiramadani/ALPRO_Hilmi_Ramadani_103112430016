package main

import (
	"fmt"
)

func main() {
	var a, b, c, max, min int
	fmt.Print("Masukkan 3 bilangan bula: ")
	fmt.Scan(&a, &b, &c)

	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	if max < c {
		max = c
	}
	if min > c {
		min = c
	}
	fmt.Println("terbesar", max)
	fmt.Println("terkecil", min)
}
