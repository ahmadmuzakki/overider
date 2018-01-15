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

	lines, attr := cleanIniLines(lines)

	dst.Attribute = attr

	digest := make(map[string]interface{})
	readline(lines, &digest)
	dst.Digest = digest
	return src, dst, nil
}

// this function will return clean lines and mapped components
func cleanIniLines(lines []string) ([]string, map[int]string) {
	newLines := make([]string, 0, len(lines))
	attr := make(map[int]string)
	for i, l := range lines {
		if l == "" || l[0] == '#' {
			attr[i] = l
			continue
		}
		newLines = append(newLines, l)
	}
	return newLines, attr
}

func (i Ini) Override(src Source, dst Destination) (Destination, error) {
	splitValue := func(line string) (string, string) {
		lines := strings.Split(line, "=")
		if len(lines) < 2 {
			return "", ""
		}
		key := lines[0]
		value := strings.Join(lines[1:], "=")
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		return key, value
	}

	for dstKey, dstVal := range dst.Digest {
		for srcKey, srcVal := range src {
			if srcKey == dstKey {
				dstLines := dstVal.([]string)
				srcLines := srcVal.([]string)

				for i, dstLine := range dstLines {
					for _, srcLine := range srcLines {

						dstLineKey, _ := splitValue(dstLine)
						srcLineKey, srcLineVal := splitValue(srcLine)
						if dstLineKey == srcLineKey {
							dstLines[i] = fmt.Sprintf("%s = %s", srcLineKey, srcLineVal)
						}
					}
				}

				dst.Digest[dstKey] = dstLines
				delete(src, dstKey)
			}
		}
	}
	return dst, nil
}

func (ini Ini) Write(dst Destination, dstURL string) error {
	raw := []string{}
	for key, val := range dst.Digest {
		raw = append(raw, key)
		raw = append(raw, val.([]string)...)
	}

	fullFile := []byte{}
	var j, i int
	attr := dst.Attribute
	for {
		data := ""
		if a, ok := attr[i]; ok {
			data = a
			delete(attr, i)
		} else if j < len(raw) {
			data = raw[j]
			j++
		}
		fullFile = append(fullFile, []byte(data+"\n")...)
		i++
		if len(attr) == 0 && len(raw) <= j {
			break
		}
	}
	fullFile = fullFile[:len(fullFile)-1] // remove the last \n

	return ioutil.WriteFile(dstURL, fullFile, 0644)
}
