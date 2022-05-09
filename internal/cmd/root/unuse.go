package root

import (
	"github.com/WAY29/hackbox/internal/cmd"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
)

func CmdUnuse(args []string) {
	prompts.Prompt.UseTool = nil
	prompts.Prompt.UseToolName = ""
}

func init() {
	suggest.AddRootSuggest("unuse", "Unuse tools")
	cmd.AddRootCommand("unuse", CmdUnuse, nil)
}
