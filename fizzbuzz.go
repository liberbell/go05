package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "Fizz Buzzs")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		}
	}
}
