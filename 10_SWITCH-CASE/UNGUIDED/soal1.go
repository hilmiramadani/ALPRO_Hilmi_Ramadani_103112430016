package main

import (
	"fmt"
)

func main() {
	var ph float32

	fmt.Print("Masukan pH air dari 0-14: ")
	fmt.Scan(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air layak minum dengan PH: ", ph)
	case ph >= 0 && ph <= 14:
		fmt.Println("Air tidak layak minum dengan PH : ", ph)
	default:
		fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")

	}
}
