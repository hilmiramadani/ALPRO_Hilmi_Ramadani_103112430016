package main

import (
	"fmt"
)

func main() {

	const kapasitas = 7

	var siswa int
	fmt.Println("Masukkan jumlah mahasiswa:")
	fmt.Scanf("%d", &siswa)

	mobil := siswa / kapasitas
	remaining := siswa % kapasitas

	if remaining > 0 {
		mobil++
		fmt.Printf("%d mobil penuh dan 1 mobil berisi %d orang\n", mobil-1, remaining)
	} else {
		fmt.Printf("%d mobil penuh\n", mobil)
	}
}
