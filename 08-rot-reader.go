package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13.r.Read(b)
	for i:= range(b) {
		b[i] = rot13(b[i])
	}
	return n, err
}

// Alternatively could add/subt depending on pos in alphabet
func rot13(c byte) byte {
	// rather not have byte to char comps (i.e., c <= 'a')
	const a, l, u = 26, 97, 65
	if c >= l && c <= l+a {
		return (c-l+13) % a + l
	}
	if c >= u && c <= u+a {
		return (c-u+13) % a + u
	}
	return c
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, r)
}
