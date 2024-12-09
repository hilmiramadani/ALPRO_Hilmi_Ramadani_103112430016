package main

import "fmt"

func main() {
	var kata string
	var jml int

	fmt.Print("masukkan kata dan jumlah:")
	fmt.Scan(&kata, &jml)

	counter := 0
	for done := false; !done; {
		fmt.Println(kata)
		counter++
		done = (counter >= jml)
	}

}
