package main

import (
	"fmt"
)

func main() {
	var bil, bagi int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bil)
	fmt.Print("Masukkan bilangan bagi: ")
	fmt.Scan(&bagi)

	jumlah := 0
	for bil >= bagi {
		bil -= bagi
		jumlah++
	}

	fmt.Print("Hasil div:", jumlah)
}
