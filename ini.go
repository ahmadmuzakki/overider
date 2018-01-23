package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Ini struct{}

var iniKeyIdentifier = regexp.MustCompile(`^\[.*\]$`)

func (i Ini) Read(srcPath, dstPath string) (Source, Destination, error) {
	src := make(map[string]interface{})
	dst := Destination{}

	readline := func(lines []string, dst *map[string]interface{}) {
		tmp := make(map[string]interface{})
		var curKey string
		for _, line := range lines {
			if iniKeyIdentifier.MatchString(line) {
				curKey = line
				tmp[curKey] = make([]string, 0)
				continue
			}
			if curKey == "" {
				continue
			}
			tmp[curKey] = append(tmp[curKey].([]string), line)
		}
		*dst = tmp
	}

	lines, err := parseFile(srcPath)
	if err != nil {
		return src, dst, err
	}
	readline(lines, &src)

	lines, err = parseFile(dstPath)
	if err != nil {
		return src, dst, err
	}

	dst = lines
	return src, dst, nil
}

func (i Ini) Override(src Source, dst Destination) (Destination, error) {
	curKey := ""
	for i, dstLine := range dst {
		if strings.TrimSpace(dstLine) == "" || dstLine[0] == '#' {
			continue
		}
		if iniKeyIdentifier.MatchString(dstLine) {
			curKey = dstLine
			continue
		}
		if curKey == "" {
			continue
		}

		parsekv := func(str string) (key, val string) {
			kv := strings.Split(str, "=")
			if len(kv) < 2 {
				return
			}

			val = strings.Join(kv[1:], "=")
			return strings.TrimSpace(kv[0]), strings.TrimSpace(val)
		}

		if srcLine, ok := src[curKey]; ok {
			dstKey, _ := parsekv(dstLine)
			for _, line := range srcLine.([]string) {
				srcKey, srcVal := parsekv(line)
				if srcKey == dstKey {
					dst[i] = fmt.Sprintf("%s = %s", dstKey, srcVal)
				}
			}
		} else {
			curKey = ""
		}
	}

	return dst, nil
}

func (ini Ini) Write(dst Destination, dstURL string) error {
	fullFile := []byte{}
	for _, line := range dst {
		fullFile = append(fullFile, []byte(line+"\n")...)
	}
	fullFile = fullFile[:len(fullFile)-1] // remove the last \n

	return ioutil.WriteFile(dstURL, fullFile, 0644)
}
