package main

import (
	"fmt"
	"os"
)

type Config struct {
	//
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	return cfg, nil
}

func main() {
	cfg, err := readConfig("config.toml")
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n", err)
	}
}
