package lexer

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Token struct {
	Name    string `yaml:"name"`
	Pattern string `yaml:"pattern"`
}

type Tokens struct {
	Tokens []Token  `yaml:"tokens"`
	Ignore []string `yaml:"ignore"`
}

func GetToken(tokenFile string) (tokens Tokens) {
	yamlData, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = yaml.Unmarshal(yamlData, &tokens)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return
	}
	return tokens
}
