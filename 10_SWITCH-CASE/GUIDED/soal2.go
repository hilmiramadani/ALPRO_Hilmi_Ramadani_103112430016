package main

import "fmt"

func main() {
	var nTanaman string
	fmt.Scanln(&nTanaman)

	switch nTanaman {
	case "nepenthes", "drosera":
		fmt.Println("termasuk tanaman karnivora")
		fmt.Println("asli indonesia")
	case "venus", "sarracenia":
		fmt.Println("termasuk tanaman karnivora")
		fmt.Println("tidak asli indonesia")
	default:
		fmt.Println("tidak termasuk tanaman karnivora")
	}
}
