package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n := len(b)
	temp_b := make([]byte, n)
	n, err := rot13.r.Read(temp_b)

	if err != nil {
		return 0, err
	}

	for i := range n {
		if db := temp_b[i]; db < 65 || db > 122 {
			b[i] = db
		} else if db < 78 || db < 110 && db > 96 {
			b[i] = db + 13
		} else {
			b[i] = db - 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
