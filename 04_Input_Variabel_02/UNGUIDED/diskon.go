package main

import "fmt"

func main() {
	var total, diskon, hasil int

	fmt.Println("total belanja:")
	fmt.Scan(&total)
	fmt.Println("diskon:")
	fmt.Scan(&diskon)

	hasil = total - (total * diskon / 100)
	fmt.Println("hasilnya:", hasil)
}
