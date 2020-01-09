package main

import "fmt"

func main() {
	x := 11
	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}
}
