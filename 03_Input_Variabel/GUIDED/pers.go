package main

import (
	"fmt"
)

func main() {

	var n, hasil float32

	fmt.Println("Masukan bilangan")
	fmt.Scanln(&n)
	hasil = (2/(n+5) + 5)
	fmt.Println("Jadi hasilnya adalah", hasil)
}
