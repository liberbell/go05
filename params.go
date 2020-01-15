package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func duble(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 0)
	fmt.Println(values)

	val := 10
	double(10)
	fmt.Println(val)
}
