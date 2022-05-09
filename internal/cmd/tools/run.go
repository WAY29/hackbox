package tools

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/WAY29/hackbox/internal/cmd"
	"github.com/WAY29/hackbox/internal/cmd/root"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/prompts/suggest"
	"github.com/WAY29/hackbox/internal/tools"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"

	"github.com/WAY29/hackbox/internal/output"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/antonmedv/expr"
)

func CmdRunSuggests(args []string, word string, currentLine string) []prompt.Suggest {
	if !strings.Contains(currentLine, "as") {
		return prompt.FilterHasPrefix(suggest.RunToolSuggests, word, true)
	} else if !strings.HasSuffix(strings.TrimSpace(currentLine), "as") {
		return prompt.FilterHasPrefix(suggest.RunToolWithOutAsSuggests, word, true)
	}

	return suggest.EmptySuggests
}

func CmdRun(args []string) {
	var (
		argStringMap = map[string]string{}
		exprValueMap = map[string]interface{}{}
		OutputBuffer bytes.Buffer
		runErr       error

		tool *tools.Tool = prompts.Prompt.UseTool
	)

	// 检查二进制文件是否存在
	if tool.BuiltinFunc == nil {
		commandArgs := strings.SplitN(prompts.Prompt.UseTool.Command, " ", 2)
		if _, err := exec.LookPath(commandArgs[0]); err != nil {
			errorMsg := fmt.Sprintf("No such executable file: %s", commandArgs[0])
			if len(prompts.Prompt.UseTool.DownloadURL) > 0 {
				errorMsg += fmt.Sprintf(". Try to download it from %s", prompts.Prompt.UseTool.DownloadURL)
			}
			Error(errorMsg)
			return
		}
	}

	saveOutputFunc := func(tool *tools.Tool, name string, outputBuffer *bytes.Buffer) {
		// 结果过滤
		data := strings.TrimSpace(outputBuffer.String())
		if len(tool.ResultFilterFunction) > 0 {
			newData := ""

			for _, filterFuncName := range strings.Split(tool.ResultFilterFunction, "|") {
				if filterFunc, ok := root.FilterFuncMap[filterFuncName]; ok {
					tempData := strings.Join(filterFunc(data), "\n")
					newData += "\n" + tempData
				}
			}

			data = strings.TrimSpace(newData)
		}

		output.Save(name, data)
	}

	background := false
	outputName := fmt.Sprintf("s%d", output.OutputNameIndex)
	hasCustomOutputName := false
	maxArgLen := len(args) - 1

	// 判断是否在后台运行以及是否自定义名字
	for i, arg := range args {
		if arg == "as" && i+1 <= maxArgLen {
			hasCustomOutputName = true
			outputName = args[i+1]
		}
		if arg == "bg" {
			background = true
		}
	}
	if !hasCustomOutputName {
		output.OutputNameIndex++
	}

	// 替换单个参数中的cmd_arg
	for _, arg := range tool.Args {
		// 对boolean做特殊处理
		name, value := arg.Name, arg.Value
		if len(value) == 0 || (arg.Type == "boolean" && (strings.ToLower(value) == "false" || value == "0")) {
			argStringMap[name] = ""
			exprValueMap[name] = false
		} else {
			if tool.BuiltinFunc != nil {
				argStringMap[name] = value
			} else {
				argStringMap[name] = strings.ReplaceAll(arg.CommandArgs, "{{}}", value)
			}
			exprValueMap[name] = true
		}
	}

	// 检查expression是否被满足
	program, err := expr.Compile(tool.ArgsExpression, expr.Env(exprValueMap))
	if err != nil {
		Error("Compile expression error: %s", err.Error())
		return
	}
	output, err := expr.Run(program, exprValueMap)
	if err != nil {
		Error("Run expression error: %s", err.Error())
		return
	}
	if exprBool, ok := output.(bool); !ok {
		Error("Expression to bool error: %s", output)
		return
	} else if !exprBool {
		Error("Unsatisfied expression: %s", tool.ArgsExpression)
		return
	}

	// 判断是否为内置工具，如果是则直接运行函数
	if tool.BuiltinFunc != nil {
		tool.BuiltinFunc(argStringMap)
		return
	}

	// 替换command args
	command := prompts.Prompt.UseTool.Command
	for name, value := range argStringMap {
		command = strings.ReplaceAll(command, fmt.Sprintf("{{%s}}", name), value)
	}

	cmd := utils.ExecCommand(command)
	Success("Run command: %s", command)

	if background {
		cmd.Stdout = &OutputBuffer
		cmd.Stderr = &OutputBuffer

		runErr = cmd.Start()
		go func() {
			cmd.Wait()
			if runErr == nil && cmd.ProcessState.ExitCode() == 0 {
				saveOutputFunc(tool, outputName, &OutputBuffer)
			}
		}()
		OutputSoonMsg(outputName)
	} else {
		multiWriter := io.MultiWriter(StandardOutput, &OutputBuffer)
		cmd.Stdout = multiWriter
		cmd.Stderr = multiWriter
		if runErr = cmd.Run(); runErr != nil {
			Error("Run command error: %s", runErr.Error())
			return
		}

		exitCode := cmd.ProcessState.ExitCode()
		if exitCode != 0 {
			Error("Run command return exit code %d", exitCode)
			return
		}

		saveOutputFunc(tool, outputName, &OutputBuffer)
		OutputMsg(outputName)
	}
}

func init() {
	suggest.AddToolSuggest("run", "Run tools")
	suggest.AddToolSuggest("go", "Run tools")
	cmd.AddToolCommand("run", CmdRun, CmdRunSuggests)
	cmd.AddToolCommand("go", CmdRun, CmdRunSuggests)
}
