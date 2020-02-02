package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
  {
    "user": "scooge McDuck",
    "type": "doposit",
    "amount": 1000000.3
    }
    `

type Request struct {
	Login  string  `json: "user"`
	Type   string  `json: "type"`
	Amount float64 `json: "amount"`
}

func main() {
	rdr := bytes.NewBufferString(data)
	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error: can`t decode - %s\n", err)
	}
	fmt.Printf("got: %+v\n", req)

	prevBalance := 8500000.0
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can`t encode - %s\n", err)
	}
}
