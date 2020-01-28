package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s - error: %s", url, err)
    return
	}

  defer resp.Body.Close()
  ctype := resp.Header.Get("content-type")
  out <- fmt.Sprintf("%s -> %s", url ctype)
}

func main() {
  urls := []string{
    "https://golang.org",
    "https://api.github.com",
    "https://httpbin.org/xml",
    "https://yahoo.co.jp",
    "https://www.ntt-at.co.jp",
  }

  ch := make(chan string)
  for _, url := range url {
    go returnType(url, ch)
  }
}
