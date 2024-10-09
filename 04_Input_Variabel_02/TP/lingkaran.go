package main

import (
	"fmt"
)

func main() {
	var r float32
	var luas float32
	var keliling float32

	fmt.Println("masukkan jari-jari lingkaran:")
	fmt.Scan(&r)

	luas = 3.14 * r * r
	keliling = 2 * 3.14 * r

	fmt.Println("hasil luas lingkaran", luas, "dan keliling lingkaran", keliling)
}
