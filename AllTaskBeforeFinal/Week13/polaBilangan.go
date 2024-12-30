package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan ganjil: ")
	fmt.Scanln(&n)

	for j := 1; j <= n; j++ {
		for i := 1; i <= n; i++ {
			if j == i || j == n-i+1 {
				fmt.Printf("%d ", j)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
