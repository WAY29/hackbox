package root

import (
	"strings"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"
)

func unique(slice1 []string, slice2 []string) []string {
	result := make([]string, 0, len(slice1)+len(slice2))
	slice1 = append(slice1, slice2...)
	temp := map[string]struct{}{}

	for _, item := range slice1 {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func CmdMergeSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) < 4 || len(args) == 5 {
		return prompt.FilterHasPrefix(*output.GetSuggests(), word, true)
	} else if len(args) == 4 {
		return prompt.FilterHasPrefix(suggest.AsSuggests, word, true)
	}

	return suggest.EmptySuggests
}

func CmdMerge(args []string) {

	if len(args) < 5 || (len(args) > 3 && args[3] != "as") {
		Error(`Invalid merge usage. e.g. merge <output1 name> <output2 name> as <output name>`)
		return
	}

	output1, output2 := args[1], args[2]
	outputName := args[4]

	if !output.IsOutputVar(output1) || !output.IsOutputVar(output2) {
		Error("Invalid output name: %s or %s", output1, output2)
		return
	}

	output1, output2 = output.Get(output1), output.Get(output2)

	if output1 == "" {
		Error("No such output: %s", output1)
		return
	}

	if output2 == "" {
		Error("No such output: %s", output2)
		return
	}

	output1Slice, output2Slice := strings.Split(output1, "\n"), strings.Split(output2, "\n")
	data := strings.Join(unique(output1Slice, output2Slice), "\n")

	output.Save(outputName, data)
	SetOutputMsg(outputName, utils.StringEllipsis(data, 70))
}

func init() {
	suggest.AddRootSuggest("merge", "Merge two outputs as one")
	cmd.AddRootCommand("merge", CmdMerge, CmdMergeSuggests)
}
