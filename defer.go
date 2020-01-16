package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("Clean up %s\n", name)
}

func worker() {
	defer cleanup("A")
	fmt.Println("worker")
}

func main() {
	worker()
}
