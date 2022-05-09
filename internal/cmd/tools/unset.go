package tools

import (
	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/c-bata/go-prompt"
)

func CmdUnSetSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) == 2 {
		return prompt.FilterHasPrefix(prompts.Prompt.UseTool.ArgSuggests, word, true)
	}
	return suggest.EmptySuggests
}

func CmdUnSet(args []string) {
	if len(args) < 2 {
		Error("Invalid unset usage. e.g. unset <arg name>")
		return
	}

	argIndex := -1

	for i, arg := range prompts.Prompt.UseTool.Args {
		if arg.Name == args[1] {
			argIndex = i
			break
		}
	}

	if argIndex == -1 {
		Error("No such argument: %s", args[1])
		return
	}

	targetArg := &prompts.Prompt.UseTool.Args[argIndex]
	targetArg.Value = ""
	SetMsg(targetArg.Name, "nil", false)
}

func init() {
	suggest.AddToolSuggest("unset", "Unset tool argument")
	cmd.AddToolCommand("unset", CmdUnSet, CmdUnSetSuggests)
}
