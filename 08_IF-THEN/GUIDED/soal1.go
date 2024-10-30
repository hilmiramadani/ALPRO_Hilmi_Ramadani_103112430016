package main

import (
	"fmt"
)

func main() {
	var n, nilai int

	fmt.Println("masukan bilangan:")
	fmt.Scan(&n)

	if n < 0 {
		nilai = -n
	} else {
		nilai = n
	}

	fmt.Println("nilai absolut dari", n, "adalah", nilai)

}
