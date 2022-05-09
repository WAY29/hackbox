package root

import (
	"github.com/WAY29/hackbox/internal/arguments"
	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/internal/tools"
	"github.com/c-bata/go-prompt"
)

func CmdUseSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) == 2 {
		return prompt.FilterHasPrefix(prompts.Prompt.ToolDir.ToolSuggests, word, true)
	}

	return suggest.EmptySuggests
}

func CmdUse(args []string) {
	if len(args) == 1 {
		Error("Invalid use usage. e.g. use <tool name>")
		return
	}

	targetTool := args[1]
	tooldir := tools.ToolDirMap[prompts.Prompt.Path]
	if tool, ok := tooldir.Tools[targetTool]; ok {
		prompts.Prompt.UseTool = tool
		prompts.Prompt.UseToolName = targetTool

		// 使用全局参数值覆盖局部参数值
		for i, arg := range tool.Args {
			if globalArg := arguments.Get(arg.Name); globalArg != nil && globalArg.Type == arg.Type {
				tool.Args[i].Value = globalArg.Value
			}
		}

	} else {
		Error("No such tools: %s", targetTool)
	}
}

func init() {
	suggest.AddRootSuggest("use", "Use tools")
	cmd.AddRootCommand("use", CmdUse, CmdUseSuggests)
}
