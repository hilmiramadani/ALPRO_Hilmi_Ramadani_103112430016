package main

import (
	"fmt"
)

func main() {
	var qirat int
	fmt.Println("Masukkan jumlah qirat: ")
	fmt.Scan(&qirat)

	fals := qirat / 6
	sisaQirat := qirat % 6
	dirham := fals / 10
	sisaFals := fals % 10
	dinar := dirham / 10
	sisaDirham := dirham % 10

	fmt.Print(dinar, sisaDirham, sisaFals, sisaQirat)
}
