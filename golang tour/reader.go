package main

import "golang.org/x/tour/reader"
import "strings"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b [] byte)(int, error){
    l := strings.NewReader("A")
	for{
	      n, err := l.Read(b)
		  return n, err
	}
}

func main() {
	reader.Validate(MyReader{})
}
