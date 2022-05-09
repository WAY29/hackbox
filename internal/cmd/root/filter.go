package root

import (
	"strings"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"
	"github.com/mingrammer/commonregex"
)

var FilterFuncMap = map[string]func(string) []string{
	"link":   commonregex.Links,
	"email":  commonregex.Emails,
	"date":   commonregex.Date,
	"time":   commonregex.Time,
	"phone":  commonregex.Phones,
	"ip":     commonregex.IPs,
	"md5":    commonregex.MD5Hexes,
	"sha1":   commonregex.SHA1Hexes,
	"sha256": commonregex.SHA256Hexes,
}

func CmdFilterSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if len(args) == 2 || len(args) == 5 {
		return prompt.FilterHasPrefix(*output.GetSuggests(), word, true)
	} else if len(args) == 3 {
		return prompt.FilterHasPrefix(suggest.FilterSuggests, word, true)
	} else if len(args) == 4 {
		return prompt.FilterHasPrefix(suggest.AsSuggests, word, true)
	}

	return suggest.EmptySuggests
}

func CmdFilter(args []string) {
	var (
		ok         bool
		filterFunc func(string) []string
		outputName string
	)

	if len(args) < 3 {
		Error(`Invalid filter usage. e.g. filter <output name> <filter> [as <new output name>]. support "|" as or condition`)
		return
	}

	inputName, filterFuncName := args[1], args[2]

	if len(args) == 5 && args[3] == "as" {
		outputName = args[4]
	}

	// 过滤器是否存在
	if filterFunc, ok = FilterFuncMap[filterFuncName]; !ok {
		Error("No such filter: %s", filterFuncName)
		return
	}

	data := output.Get(inputName)
	data = strings.Join(filterFunc(data), "\n")
	if len(data) > 0 {
		output.Save(outputName, data)
		SetOutputMsg(outputName, utils.StringEllipsis(data, 70))
	} else {
		Error("Output %s filtered value is empty", inputName)
	}
}

func init() {
	suggest.AddRootSuggest("filter", "Filter output")
	cmd.AddRootCommand("filter", CmdFilter, CmdFilterSuggests)
}
