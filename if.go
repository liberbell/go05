package main

import "fmt"

func main() {
	x := 1
	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not big")
	}
}
