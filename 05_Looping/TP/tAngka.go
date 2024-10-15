package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	target := rand.Intn(100) + 1

	var guess int
	maxAttempts := 5

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Pilihlah angka antara 1 hingga 100.")
	fmt.Println("Anda punya 5 kesempatan untuk menebak angka tersebut.")

	for attempts := 1; attempts <= maxAttempts; attempts++ {

		fmt.Printf("Percobaan ke-%d: Masukkan tebakan Anda: ", attempts)
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Terlalu kecil!")
		} else if guess > target {
			fmt.Println("Terlalu besar!")
		} else {
			fmt.Printf("Selamat! Anda berhasil menebak angka %d dengan benar!\n", target)
			return
		}
	}

	fmt.Printf("Maaf, Anda kehabisan kesempatan. Angka yang benar adalah %d.\n", target)
}
