package cmd

import (
	"github.com/c-bata/go-prompt"
)

type CmdFunc = func([]string)
type SuggestFunc = func([]string, string, string) []prompt.Suggest

var RootCmds = map[string]CmdFunc{}

var ToolCmds = map[string]CmdFunc{}

var Suggests = map[string]SuggestFunc{}

func InitCmd() {
	for key, value := range RootCmds {
		ToolCmds[key] = value
	}
}

func AddRootCommand(name string, cmdFunc CmdFunc, suggestFunc SuggestFunc) {
	RootCmds[name] = cmdFunc
	if suggestFunc != nil {
		Suggests[name] = suggestFunc
	}
}

func AddToolCommand(name string, cmdFunc CmdFunc, suggestFunc SuggestFunc) {
	ToolCmds[name] = cmdFunc
	if suggestFunc != nil {
		Suggests[name] = suggestFunc
	}
}
