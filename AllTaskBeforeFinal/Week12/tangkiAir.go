package main

import "fmt"

func main() {
	var T int
	var V int
	fmt.Scan(&T)
	VolumeSekarang := 0

	for VolumeSekarang < T {

		fmt.Scan(&V)

		VolumeSekarang += V

		if VolumeSekarang >= T {
			fmt.Println("true")
			break
		} else {
			fmt.Println("false")
		}
	}
}
