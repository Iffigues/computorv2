package lexer

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

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

func NewToken(tokenFile string) (tokens Tokens, err error) {
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
func IsValidIdentifier(s string) bool {
	r := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return r.MatchString(s)
}

func (tokens Tokens) GetToken(input string) (Token, bool) {

	for _, token := range tokens.Tokens {
		pattern, err := regexp.Compile("^" + token.Pattern)
		if err != nil {
			// Return an error if the pattern is invalid
			panic(err)
		}

		if match := pattern.FindString(input); match != "" {
			return Token{Name: token.Name, Pattern: match}, true
		}
	}

	return Token{}, false
}

func (tokens Tokens) Tokenize(input string) ([]Token, error) {
	var result []Token

	for len(input) > 0 {
		// Skip over ignored patterns
		if contains(tokens.Ignore, input) {
			input = input[1:]
			continue
		}

		// Match a token pattern using regular expressions
		token, ok := tokens.GetToken(input)
		if !ok {
			return nil, fmt.Errorf("no matching token found for input: %s", input)
		}

		result = append(result, token)
		fmt.Println(len(token.Pattern), "|", input, "|")
		input = input[len(token.Pattern):]
	}

	return result, nil
}

func contains(list []string, item string) bool {
	for _, element := range list {
		if strings.HasPrefix(item, element) {
			return strings.HasPrefix(item, element)
		}
	}
	return false
}
