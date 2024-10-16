package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("masukkan bilangan bulat 1:")
	fmt.Scan(&a)
	fmt.Println("masukkan bilangan bulat 2:")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		fmt.Print(i)
	}
}
