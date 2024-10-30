package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("masukan bilangan:")
	fmt.Scan(&n)

	if n < 0 && n%2 == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	fmt.Println()

}
