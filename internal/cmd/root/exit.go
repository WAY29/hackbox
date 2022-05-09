package root

import (
	"os"

	"github.com/WAY29/hackbox/internal/cmd"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
)

func CmdExit(args []string) {
	os.Exit(0)
}

func init() {
	suggest.AddRootSuggest("exit", "Exit hackbox")
	cmd.AddRootCommand("exit", CmdExit, nil)
}
