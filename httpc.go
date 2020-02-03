package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json: "user"`
	Action string `json: "action"`
	Count  int    `json: "count"`
}

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("Error: can`t call httpbin.org")
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
