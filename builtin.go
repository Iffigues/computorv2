package main

import (
	"fmt"
	"os"
)

var (
	builtin = map[string]func(){
		"exit": func() {
			os.Exit(0)
		},
		"help": func() {
			fmt.Println("help")
		},
	}
)
