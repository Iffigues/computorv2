package main

import (
	"computorv2/src/files"
)

func getFiles(args []string) (e []*files.Files, err error) {
	e = nil
	for _, val := range args {
		f, err := files.NewFiles(val)
		if err != nil {
			return nil, err
		} else {
			e = append(e, f)
		}
	}
	return
}
