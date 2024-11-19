package main

import (
	"fmt"
)

func main() {
	fmt.Println("Menu Restoran Cepat Saji:")
	fmt.Println("1. Burger - Rp25,000")
	fmt.Println("2. Fried Chicken - Rp20,000")
	fmt.Println("3. French Fries - Rp15,000")
	fmt.Println("4. Soft Drink - Rp10,000")
	fmt.Println("5. Coffee - Rp15,000")

	var kodeMenu int
	fmt.Print("Masukkan kode menu (1-5): ")
	fmt.Scan(&kodeMenu)

	switch kodeMenu {
	case 1:
		fmt.Println("Anda memilih Burger - Rp25,000")
	case 2:
		fmt.Println("Anda memilih Fried Chicken - Rp20,000")
	case 3:
		fmt.Println("Anda memilih French Fries - Rp15,000")
	case 4:
		fmt.Println("Anda memilih Soft Drink - Rp10,000")
	case 5:
		fmt.Println("Anda memilih Coffee - Rp15,000")
	default:
		fmt.Println("Kode menu tidak valid")
	}
}
