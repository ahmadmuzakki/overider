package main

type Overider interface {
	Read(srcPath, destPath string) (Source, Destination, error)
	Override(src Source, dst Destination) (Destination, error)
	Write(dst Destination, dstURL string) error
}

type Source map[string]interface{}

type Destination struct {
	Attribute map[int]string // attribute is a newline, or comments
	Digest    map[string]interface{}
}
