package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Inputan: ")
	fmt.Scan(&n)

	faktor := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			faktor = append(faktor, i)
			sum += i
		}
	}

	fmt.Printf("Output: ")
	if sum == n {
		fmt.Printf("Ya (karena faktor dari %d yaitu %v dan %d = %d)\n", n, faktor, sum, n)
	} else {
		fmt.Printf("Tidak (karena faktor dari %d yaitu %v dan %d ≠ %d)\n", n, faktor, sum, n)
	}
}
