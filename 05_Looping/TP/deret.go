package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	sum := 0

	fmt.Print("Deret angka dari 1 hingga ", n, ": ")
	for i := 1; i <= n; i++ {
		if i < n {
			fmt.Print(i, ", ")
		} else {
			fmt.Print(i)
		}

		sum += i
	}
	fmt.Println()

	fmt.Println("Jumlah dari 1 hingga", n, "adalah :", sum)
}
