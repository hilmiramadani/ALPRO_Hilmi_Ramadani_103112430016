package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&input)

	input = strings.ToUpper(input)

	if input == "A" || input == "I" || input == "U" || input == "E" || input == "O" {
		fmt.Println("Huruf Vokal")
	} else {
		fmt.Println("Huruf Konsonan")
	}
}
