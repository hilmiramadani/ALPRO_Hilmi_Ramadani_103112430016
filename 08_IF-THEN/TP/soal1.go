package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Nilai ujian: ")
	fmt.Scan(&nilai)

	if nilai >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
