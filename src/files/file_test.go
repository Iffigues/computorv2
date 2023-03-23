package files

import (
	"io"
	"os"
	"testing"
)

func TestFactorial(t *testing.T) {
	f, err := NewFiles("./test_files/test1")
	if err != nil {
		t.Log(err)
	}
	for {
		l, e := f.GetContent('\n')
		if e != nil {
			if e == io.EOF {
				t.Log("eof")
				break
			}
			if e == os.ErrClosed {
				t.Log("closed")
				break
			} else {
				t.Fatal(e)
			}
		}
		t.Log(l)
	}
}
