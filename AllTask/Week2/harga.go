package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Masukan harga barang: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(a/100*5 + a)
	fmt.Println(b/100*5 + b)
	fmt.Println(c/100*5 + c)
}
