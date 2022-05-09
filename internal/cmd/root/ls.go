package root

import (
	"path"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/internal/tools"
	"github.com/c-bata/go-prompt"
)

func CmdLSSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) == 2 {
		return prompt.FilterHasPrefix(prompts.Prompt.ToolDir.PathSuggests, word, true)
	}

	return suggest.EmptySuggests
}

func listToolDirectory(path string, toolDir *tools.ToolDir) {
	Text(path)

	if len(toolDir.PathSuggests) > 0 {
		Text("Subdirectory:")
		for _, suggest := range toolDir.PathSuggests {
			Text("  - %s", suggest.Text)
		}
	}

	if len(toolDir.ToolSuggests) > 0 {
		Text("Tools:")
		for _, suggest := range toolDir.ToolSuggests {
			Text("  - %s", suggest.Text)
		}
	}

	Text("")
}

func CmdLS(args []string) {
	if len(args) == 1 {
		listToolDirectory(prompts.Prompt.Path, prompts.Prompt.ToolDir)
	} else {
		targetPath := args[1]
		tempPath := prompts.Prompt.Path
		tempPath = path.Join(tempPath, targetPath)
		if tooldir, ok := tools.ToolDirMap[tempPath]; !ok {
			Error("No such tools directory: %s", tempPath)
			return
		} else {
			listToolDirectory(tempPath, tooldir)
		}
	}
}

func init() {
	suggest.AddRootSuggest("ls", "List tools and subdirectories")
	cmd.AddRootCommand("ls", CmdLS, CmdLSSuggests)
}
