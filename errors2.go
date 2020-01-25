package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	//
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can`t open configuration file")
	}
	defer file.Close()

	cfg := &Config{}
	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLoggin()
	cfg, err := readConfig("config.toml")
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n", err)
		log.Printf("Error: %+v", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
