package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
		"https://yahoo.co.jp",
		"https://www.ntt-at.co.jp",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		go func(url string) {
			wg.Add(1)
			returnType(url)
		}(url)
	}
	wg.Wait()
}
