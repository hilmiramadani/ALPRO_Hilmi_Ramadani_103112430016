package main

import "fmt"

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	var akar2 float64 = 1
	for i := 0; i <= k; i++ {
		akar2 *= float64(4*i+2) * float64(4*i+2) / (float64(4*i+1) * float64(4*i+3))
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
