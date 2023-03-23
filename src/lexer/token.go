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
	// Match a token pattern using regular expressions
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

	// Compile a regex pattern for the ignore rules
	ignorePattern, err := regexp.Compile(strings.Join(tokens.Ignore, "|"))
	if err != nil {
		return nil, err
	}

	// Loop over the input string and extract tokens
	for len(input) > 0 {
		// Skip whitespace and comments
		if match := ignorePattern.FindString(input); match != "" {
			input = input[len(match):]
			continue
		}

		// Get the next token
		token, ok := tokens.GetToken(input)
		if !ok {
			return nil, fmt.Errorf("unexpected token at position %d", len(input))
		}

		// Add the token to the result slice
		result = append(result, token)

		// Move the input cursor past the matched token
		input = input[len(token.Pattern):]
	}

	return result, nil
}
