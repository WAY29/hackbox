package root

import (
	"io/ioutil"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"
)

func CmdSetOutputSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) < 4 {
		return prompt.FilterHasPrefix(*output.GetSuggests(), word, true)
	}
	return suggest.EmptySuggests
}

func CmdSetOutput(args []string) {
	if len(args) < 3 {
		Error("Invalid setoutput usage. e.g. setoutput <output name> <output value / filepath>")
		return
	}
	name, value := args[1], args[2]

	if output.IsOutputVar(value) {
		value = output.Get(value)
	} else if utils.IsFile(value) {
		data, err := ioutil.ReadFile(value)
		if err != nil {
			Error("Read file %s error: %s", value, err)
			return
		}
		value = string(data)
	}

	output.Save(name, value)
	SetOutputMsg(name, utils.StringEllipsis(value, 70))
}

func init() {
	suggest.AddRootSuggest("setoutput", "Set output")
	cmd.AddRootCommand("setoutput", CmdSetOutput, CmdSetOutputSuggests)
}
