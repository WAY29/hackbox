package output

import (
	"strings"

	"github.com/WAY29/errors"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"
)

var OutputNameIndex = 0
var outputs = map[string]string{}
var outputSuggests = []prompt.Suggest{}

func formatName(name string) string {
	if IsOutputVar(name) {
		return name[1:]
	}
	return name
}

func IsOutputVar(name string) bool {
	return strings.HasPrefix(name, "$")
}

func Save(name string, data string) {
	name = formatName(name)

	outputs[name] = data
	setOutputSuggest(name, data)
}

func Remove(name string) bool {
	name = formatName(name)

	if _, ok := outputs[name]; ok {
		delete(outputs, name)
		removeOutputSuggest(name)
		return true
	}

	return false
}

func Get(name string) string {
	name = formatName(name)

	if data, ok := outputs[name]; ok {
		return data
	}

	return ""
}

func GetSuggests() *[]prompt.Suggest {
	return &outputSuggests
}

func GetOutputs() *map[string]string {
	return &outputs
}

func ExpandForSet(name, argType string) (string, error) {
	var result string = ""

	data := Get(name)
	if data == "" {
		return "", errors.Newf("No such output: %s", name)
	}

	// 对type类型做特殊处理
	if strings.Contains(argType, "file") {
		file, err := utils.TempFile()
		if err != nil {
			return "", errors.Wrap(err, "Create temp file failed")
		}
		file.WriteString(data)
		result = file.Name()
		file.Close()
	} else if strings.Contains(data, "\n") {
		return "", errors.Newf("Invalid %s[%s]: %s", name, argType, utils.StringEllipsis(string(data), 70))
	} else {
		result = string(data)
	}

	return result, nil
}

func setOutputSuggest(name, value string) {
	name = formatName(name)

	for i, suggest := range outputSuggests {
		if suggest.Text == "$"+name && suggest.Description != value {
			outputSuggests[i].Description = value
			return
		}
	}

	outputSuggests = append(outputSuggests, prompt.Suggest{Text: "$" + name, Description: value})
}

func removeOutputSuggest(name string) {
	name = formatName(name)

	for i, suggest := range outputSuggests {
		if suggest.Text == "$"+name {
			outputSuggests = append(outputSuggests[:i], outputSuggests[i+1:]...)
		}
	}
}
