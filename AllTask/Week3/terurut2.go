package main

import "fmt"

func main() {
	var (
		bilangan, angka1, angka2, angka3 int
	)

	fmt.Print("Input bilangan 3 digit positif: ")
	fmt.Scan(&bilangan)

	angka1 = bilangan / 100
	angka2 = (bilangan % 100) / 10
	angka3 = bilangan % 10

	if angka1 < angka2 && angka2 < angka3 {
		fmt.Println("true")
		fmt.Println("terurut secara membesar", bilangan)
	} else if angka1 > angka2 && angka2 > angka3 {
		fmt.Println("true")
		fmt.Println("terurut mengecil", bilangan)
	} else {
		fmt.Println("false")
		fmt.Println("tidak terurut", bilangan)
	}
}
