package lexer

import (
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

func GetToken(tokenFile string) (tokens Tokens, err error) {
	yamlData, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return tokens, err
	}

	err = yaml.Unmarshal(yamlData, &tokens)
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}
