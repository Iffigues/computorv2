package main

type vars map[string]*string
type funcs map[string]*string

type context struct {
	vars  vars
	funcs funcs
}

func newContext() (f *context) {
	return &context{
		vars:  make(vars),
		funcs: make(funcs),
	}
}
