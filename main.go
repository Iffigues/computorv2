package main

import (
	"computorv2/src/lexer"
	"flag"
	"fmt"
	"os"
)

func setFlagHelp() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s <[OPTIONS]>\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "\n Options:")
		fmt.Fprintln(flag.CommandLine.Output(), "  -h, --help  display this help message")
		fmt.Fprintln(flag.CommandLine.Output(), "  -history    change prompt history default: /tmp/readline.tmp")
		fmt.Fprintln(flag.CommandLine.Output(), "  <arg>: load file to execute calcul, work with histrory file")
	}
}

func main() {
	lex, err := lexer.NewToken("./data/token.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	history := flag.String("history", "/tmp/readline.tmp", "log file history")

	setFlagHelp()

	flag.Parse()
	args := flag.Args()
	f, err := getFiles(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range f {
		fmt.Println(val)
	}

	if err := terms("@>:", *history, lex); err != nil {
		fmt.Println(err)
	}
}
