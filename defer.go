package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("Clean up %s\n", name)
}

func main() {
	worker()
}
