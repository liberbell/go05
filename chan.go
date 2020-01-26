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

	fmt.Println("--------")

	go func() {
		_for i := 0; i < 3; i++ {
			fmt.Printf("sendign: %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()
}
