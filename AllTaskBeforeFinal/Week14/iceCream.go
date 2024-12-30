package main

import (
	"fmt"
)

func main() {
	var kTidur, tMain, tKebun int
	var kondisi1, kondisi2, kondisi3 bool
	var hasil string

	fmt.Scan(&kTidur, &tMain, &tKebun)

	kondisi3 = kTidur >= 60 && tMain >= 75 && tKebun >= 60
	kondisi2 = kTidur >= 80 && tKebun >= 80
	kondisi1 = kTidur >= 100

	if kondisi1 || kondisi2 || kondisi3 {
		hasil = "Ice Cream"
	} else {
		hasil = "Tidak"
	}

	fmt.Println(hasil)
}
