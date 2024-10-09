package main

import "fmt"

func main() {
	var bilangan int

	fmt.Println("masukkan bilangan 3 digit:")
	fmt.Scan(&bilangan)

	d1 := bilangan / 100
	d2 := (bilangan / 10) % 10
	d3 := bilangan % 10

	if d1 < d2 && d2 < d3 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
