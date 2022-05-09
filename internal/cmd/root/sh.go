package root

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/WAY29/hackbox/internal/cmd"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/utils"

	"github.com/c-bata/go-prompt"
)

func CmdSHSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	return prompt.FilterHasPrefix(*output.GetSuggests(), word, true)
}

func CmdSH(args []string) {
	var OutputBuffer bytes.Buffer

	if len(args) < 2 {
		Error("Invalid sh usage. e.g. sh <command>")
		return
	}

	command := strings.Join(args[1:], " ")
	outputs := *output.GetOutputs()
	envs := make([]string, 0, len(outputs))
	for name, value := range outputs {
		envs = append(envs, name+"="+value)
	}
	envs = append(envs, os.Environ()...)

	cmd := utils.ExecCommand(command)
	Success("Run command: %s", command)

	multiWriter := io.MultiWriter(StandardOutput, &OutputBuffer)
	cmd.Stdin = os.Stdin
	cmd.Stdout = multiWriter
	cmd.Stderr = multiWriter
	cmd.Env = envs

	if err := cmd.Run(); err != nil {
		Error("Run command %s error: %s", command, err.Error())
		return
	}

	outputName := fmt.Sprintf("s%d", output.OutputNameIndex)
	output.OutputNameIndex++

	output.Save(outputName, OutputBuffer.String())
	OutputMsg(outputName)
}

func init() {
	suggest.AddRootSuggest("sh", "Run local shell command")
	cmd.AddRootCommand("sh", CmdSH, CmdSHSuggests)
}
