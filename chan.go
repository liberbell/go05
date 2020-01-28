package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// <-ch
	// fmt.Println("Here")

	go func() {
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("Got: %d\n", val)

	fmt.Println("--------")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sendign: %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("Recieved: %d\n", val)
	}
}
