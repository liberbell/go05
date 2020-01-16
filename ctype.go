package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, err) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}
}
