package main

import (
	"fmt"
)

func main() {
	var input string

	for {
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&input)

		if input == "telkom" {
			fmt.Println("Program selesai.")
			break
		} else {
			fmt.Printf("Anda mengetik: %s\n", input)
		}
	}
}
