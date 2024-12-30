package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scanln(&a, &b, &c)

	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Sisi segitiga tidak valid. Sisi harus bilangan positif.")
	} else if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == b || b == c || a == c {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}
