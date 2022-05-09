package root

import (
	"fmt"
	"text/tabwriter"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"
)

func CmdOutputSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) == 2 {
		return prompt.FilterHasPrefix(*output.GetSuggests(), word, true)
	}
	return suggest.EmptySuggests
}

func CmdOutput(args []string) {
	var data string

	if len(args) == 1 {
		w := tabwriter.NewWriter(StandardOutput, 10, 20, 1, ' ', 0)
		fmt.Fprintln(w, "OutputName\tData")
		for name, outputByte := range *output.GetOutputs() {
			fmt.Fprintf(w, "%s\t%s\n", name, utils.StringEllipsis(string(outputByte), 70))
		}
		fmt.Fprintln(w, "")

		w.Flush()
		return
	}

	data = output.Get(args[1])
	if len(data) > 0 {
		Text(data + "\n")
	}

}

func init() {
	suggest.AddRootSuggest("output", "Show output")
	cmd.AddRootCommand("output", CmdOutput, CmdOutputSuggests)
}
