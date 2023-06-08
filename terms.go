package main

import (
	"computorv2/src/lexer"
	"computorv2/src/prompt"
	"fmt"
)

func isBuiltin(a string) bool {
	for key, val := range builtin {
		if key == a {
			val()
			return true
		}
	}
	return false
}

func execToken(lex lexer.Tokens, a string) {
	token, err := lex.Tokenize(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(token)
	}
}

func terms(a, h string, lex lexer.Tokens) (err error) {
	p, err := prompt.GetPrompt(a, h)
	if err != nil {
		return err
	}
	defer p.Closes()

	var text string
	for {
		text, err = p.Line()
		if err != nil {
			break
		}
		if !isBuiltin(text) {
			execToken(lex, text)
		}
	}
	return nil
}
