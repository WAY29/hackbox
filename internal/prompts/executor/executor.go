package executor

import (
	"strings"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/prompts"

	"github.com/google/shlex"
)

func Executor(in string) {
	var funcs map[string]func([]string)

	// ! to run local sh
	if strings.HasPrefix(in, "!") {
		command := in[1:]
		cmd.RootCmds["sh"]([]string{"sh", command})
		return
	}

	args, err := shlex.Split(in)
	if err != nil {
		Text("Invalid command: %s", in)
	}

	if len(args) == 0 {
		return
	}

	command := args[0]

	if prompts.Prompt.UseTool == nil {
		funcs = cmd.RootCmds
	} else {
		funcs = cmd.ToolCmds
	}

	if commandFunc, ok := funcs[command]; ok {
		commandFunc(args)
	} else {
		Error("No such command: %s", command)
	}
}
