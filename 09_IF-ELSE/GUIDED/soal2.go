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
	if input < "A" || input > "Z" {
		fmt.Println("Bukan huruf")
	} else if input == "A" || input == "I" || input == "U" || input == "E" || input == "O" {
		fmt.Println("Vokal")
	} else {
		fmt.Println("konsonan")
	}
}
