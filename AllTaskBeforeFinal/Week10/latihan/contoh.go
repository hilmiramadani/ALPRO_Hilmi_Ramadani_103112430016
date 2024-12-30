package main

import "fmt"

func main() {
	var val int

	fmt.Scan(&val)

	if val >= 0 {
		fmt.Println("positive")
		if val%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	} else {
		fmt.Print("Number is")
		fmt.Print("Negative")
	}
	fmt.Println("finish")
}
