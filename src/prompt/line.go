package prompt


import (
	"github.com/chzyer/readline"
)

func confLine(a , b string) *readline.Config {
	return  &readline.Config{
		Prompt:          a,
		HistoryFile:     b,
		HistorySearchFold:   true,
	}
}
