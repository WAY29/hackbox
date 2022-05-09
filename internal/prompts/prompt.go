package prompts

import (
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/internal/tools"
)

type promptOption struct {
	Path        string
	UseToolName string
	UseTool     *tools.Tool
	ToolDir     *tools.ToolDir
}

var (
	Prompt promptOption
)

func InitPrompt(path string) {
	Prompt = promptOption{
		Path: path,
	}
	Prompt.ToolDir = tools.ToolDirMap[path]
	suggest.InitSuggests()
}
