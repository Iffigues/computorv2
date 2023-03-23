package files

import (
	"bufio"
	"io"
	"os"
)

type Files struct {
	file   *os.File
	reader *bufio.Reader
}

func NewFiles(f string) (s *Files, err error) {
	s = new(Files)
	s.file, err = os.Open(f)
	if err != nil {
		return nil, err
	}
	s.reader = bufio.NewReader(s.file)
	return
}

func (f *Files) Close() (err error) {
	return f.file.Close()
}

func (f *Files) GetContent(e byte) (s string, err error) {
	ss, err := f.reader.ReadBytes(e)
	return string(ss), err
}

func (f *Files) GetAll(e byte) (s []string, err error) {
	for {
		line, err := f.GetContent(e)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		s = append(s, line)
	}
	return
}
