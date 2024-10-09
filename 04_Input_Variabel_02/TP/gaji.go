package main

import "fmt"

func main() {
	var jKerjaMingguan float64
	var upahPerJam float64
	var jNormal, jLembur, totalGajiMingguan float64

	fmt.Println("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scan(&jKerjaMingguan)

	fmt.Println("Masukkan upah per jam: ")
	fmt.Scan(&upahPerJam)

	if jKerjaMingguan > 40 {
		jNormal = 40
		jLembur = jKerjaMingguan - 40
	} else {
		jNormal = jKerjaMingguan
		jLembur = 0
	}

	totalGajiMingguan = (jNormal * upahPerJam) + (jLembur * 1.5 * upahPerJam)

	totalGajiBulanan := totalGajiMingguan * 4

	fmt.Println("Total gaji bulanan: Rp", int(totalGajiBulanan))

}
