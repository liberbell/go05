package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://httpbin.og/get")
	if err != nil {
		log.Fatalf("Error: can`t call httpbin.org")
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
