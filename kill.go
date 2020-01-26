package main

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "can`t open pid file (is server runnig?)")
	}
}
