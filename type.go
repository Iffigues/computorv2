package main

import (
	"computorv2/src/lexer"
)

type vars map[string]*string
type funcs map[string]*string

type context struct {
	vars  vars
	funcs funcs
}

type computor struct {
	lexer lexer.Tokens
}

func newContext() (f *context) {
	return &context{
		vars:  make(vars),
		funcs: make(funcs),
	}
}
