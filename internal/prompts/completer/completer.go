package completer

import (
	"strings"

	"github.com/WAY29/hackbox/internal/cmd"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/c-bata/go-prompt"
)

func Completer(in prompt.Document) []prompt.Suggest {
	currentLine := in.CurrentLine()
	w := in.GetWordBeforeCursor()
	args := strings.Split(currentLine, " ")

	if len(args) == 0 {
		return suggest.EmptySuggests
	}

	if len(args) == 1 {
		if prompts.Prompt.UseTool != nil {
			return prompt.FilterHasPrefix(suggest.ToolSuggests, args[0], true)

		} else {
			return prompt.FilterHasPrefix(suggest.RootSuggests, args[0], true)
		}
	}

	if suggestFunc, ok := cmd.Suggests[args[0]]; ok {
		return suggestFunc(args, w, currentLine)
	}

	return suggest.EmptySuggests
}
