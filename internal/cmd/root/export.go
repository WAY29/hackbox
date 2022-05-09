package root

import (
	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"
)

func CmdExportSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) == 2 {
		return prompt.FilterHasPrefix(*output.GetSuggests(), word, true)
	}
	return suggest.EmptySuggests
}

func CmdExport(args []string) {
	var (
		path string
		name string
	)
	if len(args) < 2 {
		Error("Invalid export usage. e.g. export <output name> [export path, default <output name>.txt]")
		return
	}

	name = args[1]

	if len(args) < 3 {
		path = name + ".txt"
	} else {
		path = args[2]
	}

	data := output.Get(name)
	if data != "" {
		utils.WriteFile(path, []byte(data))
		ExportMsg(name, path)
	} else {
		Error("No such output: %s", name)
	}
}

func init() {
	suggest.AddRootSuggest("export", "Export output as file")
	cmd.AddRootCommand("export", CmdExport, CmdExportSuggests)
}
