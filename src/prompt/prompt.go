package prompt

import (
	"strings"
	"github.com/chzyer/readline"
)

func GetPrompt(a, b string) (e *prompt, err error) {
	rl, err := readline.NewEx(confLine(a, b))
	if err != nil {
		return nil, err
	}
	return &prompt{
		prompt: rl,
	}, nil
}

func (p *prompt) Closes() {
	p.prompt.Close()
}

func (p *prompt) read() (text string, err error) {
	return p.prompt.Readline()
}

func (p *prompt) Line() (text string, err error) {
	text, err = p.read()
	if err != nil {
		return
	}
	return strings.TrimSpace(text), nil
}

