package main

import (
	"computorv2/src/prompt"
	"fmt"
)

func terms(a, h string) (err error) {
	p, err := prompt.GetPrompt(a, h)
	if err != nil {
		return err
	}
	defer p.Closes()

	var text string
	for {
		text, err = p.Line()
		if err != nil {
			fmt.Println(err)
			err = nil
			break
		}
		println(text)
	}
	return nil
}
