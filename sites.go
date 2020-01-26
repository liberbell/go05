package main

import "net/http"

func returnType(url string) {
	resp, err := http.Get(url)
}
