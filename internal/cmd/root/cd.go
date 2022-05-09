package root

import (
	"path"
	"strings"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/internal/tools"

	"github.com/c-bata/go-prompt"
)

func GetTargetPath(inputPath string) string {
	targetPath := prompts.Prompt.Path

	if strings.HasPrefix(inputPath, "/") {
		targetPath = inputPath
	} else if len(inputPath) > 0 {
		targetPath = path.Join(targetPath, inputPath)
	} else {
		return ""
	}

	return targetPath
}

func CmdCDSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	var pathStringBuilder strings.Builder

	if len(args) == 2 {
		tempPath := word[:strings.LastIndex(word, "/")+1]
		targetPath := GetTargetPath(tempPath)

		if tooldir, ok := tools.ToolDirMap[targetPath]; ok {
			// temporary change tooldir.PathSuggests
			restoreSuggests := make([]string, 0, len(tooldir.PathSuggests))
			for _, suggest := range tooldir.PathSuggests {
				restoreSuggests = append(restoreSuggests, suggest.Text)
			}
			for i := range tooldir.PathSuggests {
				pathStringBuilder.WriteString(tempPath)
				if !strings.HasSuffix(tempPath, "/") {
					pathStringBuilder.WriteString("/")
				}
				pathStringBuilder.WriteString(tooldir.PathSuggests[i].Text)

				tooldir.PathSuggests[i].Text = pathStringBuilder.String()

				pathStringBuilder.Reset()
			}

			defer func() {
				for i := range tooldir.PathSuggests {
					tooldir.PathSuggests[i].Text = restoreSuggests[i]
				}
			}()

			return prompt.FilterHasPrefix(tooldir.PathSuggests, word, true)
		} else {
			return prompt.FilterHasPrefix(prompts.Prompt.ToolDir.PathSuggests, word, true)
		}
	}

	return suggest.EmptySuggests
}

func CmdCD(args []string) {
	if len(args) == 1 {
		Error("Invalid cd usage. e.g. cd <tool directory>")
		return
	}

	targetPath := GetTargetPath(args[1])

	if tooldir, ok := tools.ToolDirMap[targetPath]; !ok {
		Error("No such tools directory: %s", targetPath)
		return
	} else {
		prompts.Prompt.ToolDir = tooldir
		prompts.Prompt.Path = targetPath
		prompts.Prompt.UseTool = nil
		prompts.Prompt.UseToolName = ""
	}
}

func init() {
	suggest.AddRootSuggest("cd", "Change tool directory")
	cmd.AddRootCommand("cd", CmdCD, CmdCDSuggests)
}
