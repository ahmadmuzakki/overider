package main

import (
	"io/ioutil"
	"strings"
)

func parseFile(path string) ([]string, error) {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(byt), "\n")
	return lines, nil
}
