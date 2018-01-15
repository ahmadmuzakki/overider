package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	if len(os.Args) < 2 {
		return 1
	}

	source := os.Args[1]
	destination := os.Args[2]

	log.Println(source, destination)
	ini := Ini{}
	src, dst, err := ini.Read(source, destination)
	if err != nil {
		log.Println(err)
		return 1
	}

	dst, err = ini.Override(src, dst)
	if err != nil {
		log.Println(err)
		return 1
	}

	if os.Args[3] == "-o" && os.Args[4] != "" {
		destination = os.Args[4]
	}

	err = ini.Write(dst, destination)
	if err != nil {
		log.Println(err)
		return 1
	}
	return 0
}
