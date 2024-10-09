package main

import "fmt"

func main() {
	var berat, tinggi float32

	fmt.Println("masukkan berat dan tinggi:")
	fmt.Scan(&berat, &tinggi)

	bmi := berat / (tinggi * tinggi)
	fmt.Println(bmi)
}
