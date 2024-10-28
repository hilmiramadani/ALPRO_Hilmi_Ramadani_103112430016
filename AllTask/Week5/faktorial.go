package main

import "fmt"

func main() {
	var n int
	var faktorial int = 1

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Bilangan harus positif!")
		return
	}

	for i := 1; i <= n; i++ {
		faktorial *= i
	}

	fmt.Printf("%d = %d\n", n, faktorial)
}
