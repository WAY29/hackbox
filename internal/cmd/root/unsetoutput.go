package root

import (
	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/c-bata/go-prompt"
)

func CmdUnSetOutputSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) == 2 {
		return prompt.FilterHasPrefix(*output.GetSuggests(), word, true)
	}
	return suggest.EmptySuggests
}

func CmdUnSetOutput(args []string) {
	if len(args) < 2 {
		Error("Invalid unsetoutput usage. e.g. unsetoutput <output name>")
		return
	}
	name := args[1]

	removed := output.Remove(name)
	if removed {
		SetOutputMsg(name, "nil")
	} else {
		Error("No such output: %s", name)
	}
}

func init() {
	suggest.AddRootSuggest("unsetoutput", "UnSet output")
	cmd.AddRootCommand("unsetoutput", CmdUnSetOutput, CmdUnSetOutputSuggests)
}
