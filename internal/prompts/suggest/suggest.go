package suggest

import "github.com/c-bata/go-prompt"

var (
	BackGroundSuggest = prompt.Suggest{Text: "bg", Description: "Run tools in background"}
	AsSuggest         = prompt.Suggest{Text: "as", Description: "Set the output name for later use"}
)

var (
	EmptySuggests = []prompt.Suggest{}
	RootSuggests  = []prompt.Suggest{}
	ToolSuggests  = []prompt.Suggest{}

	RunToolSuggests = []prompt.Suggest{
		BackGroundSuggest,
		AsSuggest,
	}
	RunToolWithOutAsSuggests = []prompt.Suggest{
		BackGroundSuggest,
	}
	AsSuggests = []prompt.Suggest{
		AsSuggest,
	}

	FilterSuggests = []prompt.Suggest{
		{Text: "link", Description: "Filter link"},
		{Text: "email", Description: "Filter email"},
		{Text: "date", Description: "Filter date"},
		{Text: "time", Description: "Filter time"},
		{Text: "phone", Description: "Filter phone"},
		{Text: "ip", Description: "Filter ip"},
		{Text: "md5", Description: "Filter md5"},
		{Text: "sha1", Description: "Filter sha1"},
		{Text: "sha256", Description: "Filter sha256"},
	}
)

func InitSuggests() {
	ToolSuggests = append(ToolSuggests, RootSuggests...)
}

func AddRootSuggest(text, descriptions string) {
	RootSuggests = append(RootSuggests, prompt.Suggest{Text: text, Description: descriptions})
}

func AddToolSuggest(text, descriptions string) {
	ToolSuggests = append(ToolSuggests, prompt.Suggest{Text: text, Description: descriptions})
}
