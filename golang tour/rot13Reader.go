package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func (p *rot13Reader) Read(b [] byte) (n int ,  err error){
    b = make([] byte, 8)
	n, err = p.r.Read(b)
	for i:= 0; i< 8; i++{
	   b[i]  = b[i] +13
	}
	return
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
