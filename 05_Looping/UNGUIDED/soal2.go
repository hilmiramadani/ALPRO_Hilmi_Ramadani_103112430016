package main

import "fmt"

func main() {
	var j, a float32
	var n int

	fmt.Println("masukkan n")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println("masukkan jari jari:")
		fmt.Scan(&j)
		fmt.Println("masukkan alas:")
		fmt.Scan(&a)
		rumus := 3.14 * j * j * a * (1.0 / 3.0)
		fmt.Println("Volumnya:", rumus)
	}

}
