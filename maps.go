package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	fmt.Println(len(stocks))
}
