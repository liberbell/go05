package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("Clean up %s\n", name)
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("worker")
}

func main() {
	worker()
}
