package tools

import (
	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/internal/tools"
	"github.com/c-bata/go-prompt"
)

func CmdShowSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) == 2 {
		return prompt.FilterHasPrefix(prompts.Prompt.UseTool.ArgSuggests, word, true)
	}
	return suggest.EmptySuggests
}

func CmdShow(args []string) {
	if prompts.Prompt.UseTool == nil {
		return
	}

	if len(args) == 1 {
		Text("Arguments:")
		for _, arg := range prompts.Prompt.UseTool.Args {
			ArgumentMsg(arg.Name, arg.Type, arg.Value)
		}
		return
	}

	var targetArg *tools.Arg
	tool := prompts.Prompt.UseTool

	for i, arg := range tool.Args {
		if arg.Name == args[1] {
			targetArg = &tool.Args[i]
			break
		}
	}

	if targetArg == nil {
		Error("No such argument: %s", args[1])
		return
	}

	ArgumentMsg(targetArg.Name, targetArg.Type, targetArg.Value)
}

func init() {
	suggest.AddToolSuggest("show", "Show tool argument(s)")
	cmd.AddToolCommand("show", CmdShow, CmdShowSuggests)
}
