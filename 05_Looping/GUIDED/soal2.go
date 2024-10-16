package main

import "fmt"

func main() {
	var a, t, n int

	fmt.Println("masukkan n")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println("masukkan alas dan tinggi:")
		fmt.Scan(&a, &t)
		rumus := a * t / 2
		fmt.Println("Luasnya:", rumus)
	}

}
