package main

import "fmt"

func main() {
	var bil int

	fmt.Println("masukkan bilangan bulat positif:")
	for {
		fmt.Scan(&bil)

		if bil <= 0 {
			bil = 0
			continue
		}
		break
	}

	fmt.Printf("bilangan bulat positif: %d\n", bil)

}
