package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	os.Exit(Main())
}

func Main() int {

	lenargs := len(os.Args)
	if lenargs < 2 {
		return 1
	}

	source := os.Args[1]
	destination := os.Args[2]

	ovr := getOverider(source, destination)

	src, dst, err := ovr.Read(source, destination)
	if err != nil {
		log.Println(err)
		return 1
	}

	dst, err = ovr.Override(src, dst)
	if err != nil {
		log.Println(err)
		return 1
	}

	if lenargs == 5 && os.Args[3] == "-o" && os.Args[4] != "" {
		destination = os.Args[4]
	}

	err = ovr.Write(dst, destination)
	if err != nil {
		log.Println(err)
		return 1
	}
	return 0
}

func getOverider(source, destination string) Overider {
	srcSplit := strings.Split(source, ".")
	dstSplit := strings.Split(destination, ".")

	ext1 := srcSplit[len(srcSplit)-1]
	ext2 := dstSplit[len(dstSplit)-1]
	if ext1 != ext2 {
		log.Fatal("Source file and Destination file is different type")
	}

	if ext1 == "ini" {
		return Ini{}
	}

	log.Fatal(fmt.Sprintf("%s file is not supported", ext1))
	return nil
}
