package main

import "fmt"

func main() {
	const maxAttempts = 3
	var password string
	var correctPassword = "12345"
	var attempts int

	fmt.Println("Selamat datang!")

	for attempts < maxAttempts {
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if password == correctPassword {
			fmt.Println("Login berhasil!")
			return
		} else {
			fmt.Println("Password salah. Silakan coba lagi")
			attempts++
		}
	}

	fmt.Println("Login ditolak.")
}
