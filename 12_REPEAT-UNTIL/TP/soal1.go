package main

import (
	"fmt"
)

func main() {
	angkaRahasia := 3
	var tebakan int

	for {
		fmt.Print("Tebak angka (1-10):")
		fmt.Scan(&tebakan)

		if tebakan == angkaRahasia {
			fmt.Println("Selamat! tebakan Anda benar.")
			break
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}
