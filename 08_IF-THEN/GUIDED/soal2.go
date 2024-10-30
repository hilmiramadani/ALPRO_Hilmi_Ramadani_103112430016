package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("masukan bilangan:")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("bukan positif")
	} else {
		fmt.Println("positif")
	}
	fmt.Println()

}
