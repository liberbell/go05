package main

import "fmt"

func main() {
	ch := make(chan int)

	// <-ch
	// fmt.Println("Here")

	go func() {
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("Got: %d\n", val)
}
