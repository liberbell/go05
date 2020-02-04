package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Gophers.")
}

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
