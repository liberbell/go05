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
