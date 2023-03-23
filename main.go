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
	lexer.GetToken("./token.yaml")
	history := flag.String("history", "/tmp/readline.tmp", "log file history")

	setFlagHelp()

	flag.Parse()
	args := flag.Args()
	f, err := getFiles(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if f != nil {

	}

	if err := terms("@>:", *history); err != nil {
		fmt.Println(err)
	}
}
