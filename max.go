package main

import "fmt"

func main() {
	nums := []int{18, 29, 3, 43, 7, 13, 119}

	max := nums[0]
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)
}
