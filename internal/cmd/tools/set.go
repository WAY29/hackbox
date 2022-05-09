package tools

import (
	"fmt"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func CmdSetSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if prompts.Prompt.UseTool == nil {
		return suggest.EmptySuggests
	}

	if len(args) == 3 {
		return prompt.FilterHasPrefix(*output.GetSuggests(), word, true)
	} else if len(args) == 2 {
		return prompt.FilterHasPrefix(prompts.Prompt.UseTool.ArgSuggests, word, true)
	}

	return suggest.EmptySuggests
}

func CmdSet(args []string) {
	var err error

	if len(args) < 3 {
		Error("Invalid set usage. e.g. set <arg name> <arg value>")
		return
	}

	value := args[2]
	argIndex := -1

	// 寻找arg是否存在
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

	// 对output变量进行处理
	if output.IsOutputVar(value) {
		value, err = output.ExpandForSet(value, targetArg.Type)
		if err != nil {
			Error(err.Error())
			return
		}
	}

	if targetArg.Type != "string" {
		err := validate.Var(value, fmt.Sprintf("required,%s", targetArg.Type))
		if err != nil {
			Error("Invalid %s[%s]: %s", targetArg.Name, targetArg.Type, utils.StringEllipsis(value, 70))
			return
		}
	}

	targetArg.Value = value
	SetMsg(targetArg.Name, value, false)
}

func init() {
	suggest.AddToolSuggest("set", "Set tool argument")
	cmd.AddToolCommand("set", CmdSet, CmdSetSuggests)
}
