package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func rot13(x byte) byte {
	if (x>='A' && x<'N' || x>='a' && x<'n'){
		x=  x+13
	}else{
		if (x!=' '){
			x=  x-13
		}
	}
	return x
}

func (self rot13Reader) Read(slice []byte) (n int,e error){
	n,err:= self.reader.Read(slice)
	for i:=0; i < n; i++{
		slice[i] = rot13(slice[i])
	}
	return n,err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
