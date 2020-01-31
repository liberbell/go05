package main

import (
	"log"
	"os"
)

type Config struct {
	Login struct {
		User     string
		Password string
	}
}

func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("error: can`t open config file - %s\n", err)
	}
}
