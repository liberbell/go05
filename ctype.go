package main

import "net/http"

func contentType(url string) (string, err) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
}
