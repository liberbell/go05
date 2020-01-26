package main

import "io/ioutil"

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
}
