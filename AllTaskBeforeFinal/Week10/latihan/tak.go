package main

import "fmt"

func main() {
	var nilai float32

	fmt.Scanln(&nilai)
	if nilai < 2.00 {
		fmt.Println("Poor")
	} else if nilai >= 2.00 && nilai <= 2.75 {
		fmt.Println("Fair")
	} else if nilai >= 2.76 && nilai <= 3.00 {
		fmt.Println("Satifactory")
	} else if nilai >= 3.01 && nilai <= 3.50 {
		fmt.Println("Very Good")
	} else {
		fmt.Println("Excellent")
	}
}
