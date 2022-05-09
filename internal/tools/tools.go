package tools

import "github.com/c-bata/go-prompt"

type BuiltinFunc = func(map[string]string)

type Tool struct {
	Descriptions         string `yaml:"descriptions"`
	DownloadURL          string `yaml:"download_url"`
	Args                 []Arg  `yaml:"args"`
	Command              string `yaml:"command"`
	ResultFilterFunction string `yaml:"result_filter_function"`
	ArgsExpression       string `yaml:"args_expression"`
	BuiltinFunc          BuiltinFunc
	ArgSuggests          []prompt.Suggest
}

type Arg struct {
	Name         string `yaml:"name"`
	Type         string `yaml:"type"`
	Descriptions string `yaml:"descriptions"`
	CommandArgs  string `yaml:"cmd_arg"`
	Value        string
}

type Tools = map[string]*Tool

type Conf struct {
	Tools Tools `yaml:"tools"`
}

type ToolDir struct {
	Tools        Tools
	Children     map[string]*ToolDir
	PathSuggests []prompt.Suggest
	ToolSuggests []prompt.Suggest
}

var ToolDirMap = map[string]*ToolDir{
	"/": {
		Tools:    make(map[string]*Tool),
		Children: make(map[string]*ToolDir),
	},
}
