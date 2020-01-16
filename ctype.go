package main

func contentType(url string) (string, err) {
  resp, err := http.Get(url)
  if err != nil {
    retrun "", err
  }
}
