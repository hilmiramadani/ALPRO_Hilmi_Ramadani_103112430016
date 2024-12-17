package main

import "fmt"

func main() {
	var n int

	fmt.Println("masukkan bilangan bulat:")
	fmt.Scan(&n)

	fmt.Printf("faktor dari %d adalah ", n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
