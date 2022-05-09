package tools

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/utils"
	"github.com/c-bata/go-prompt"

	"github.com/WAY29/errors"
	_ "github.com/WAY29/errors"
	"gopkg.in/yaml.v2"
)

func GetToolPath(customPath string) string {
	if customPath != "" {
		return customPath
	}

	// 从当前目录中寻找配置
	tempToolPath := "./tools.yaml"
	if utils.IsFile(tempToolPath) {
		return tempToolPath
	}

	// 从家目录中寻找配置
	homeDir, err := utils.HomeDir()
	if err != nil {
		Error("Find Home directory error")
		return ""
	}

	tempToolConfigPath := path.Join(homeDir, ".config", "hackbox")
	tempToolPath = path.Join(tempToolConfigPath, "tools.yaml")
	// 如果没有配置则创建一份默认配置
	if !utils.IsFile(tempToolPath) {
		os.MkdirAll(tempToolConfigPath, 0755)
		if err := utils.WriteFile(tempToolPath, defaultToolsContent); err != nil {
			Error("Write %s error: %s", tempToolPath, err.Error())
		}
	}

	return tempToolPath
}

func LoadTools(toolPath string, quiet bool) error {
	// 读取文件
	data, err := ioutil.ReadFile(toolPath)
	if err != nil {
		return errors.Wrapf(err, "Load tools[%s] error", toolPath)
	}
	// 解析yaml
	conf := Conf{}
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		return errors.Wrapf(err, "Parse tools[%s] error", toolPath)
	}

	// 配置tools
	ConfigTools(conf.Tools)

	// 配置tooldirs
	ConfigToolDirs(ToolDirMap)

	if !quiet {
		Success("Load tools from %s", toolPath)
	}

	return nil
}

func ConfigTools(tools Tools) {

	// 遍历配置tool
	for name, tool := range tools {
		ConfigTool(name, tool, false)
	}
}

func ConfigTool(name string, tool *Tool, immediatelyConfigToolDir bool) {
	var (
		newToolDir      *ToolDir
		lastToolDir     *ToolDir        = ToolDirMap["/"]
		tempPathBuilder strings.Builder = strings.Builder{}
	)

	NameSplitSlice := strings.Split(name, ".")
	maxIndex := len(NameSplitSlice) - 1

	for i, s := range NameSplitSlice {
		// 为最后的ToolDir写入Tool
		if i == maxIndex {
			lastToolDir.Tools[s] = tool
			break
		}

		tempPathBuilder.WriteString("/")
		tempPathBuilder.WriteString(s)
		if tooldir, ok := ToolDirMap[tempPathBuilder.String()]; !ok {
			newToolDir = &ToolDir{
				Tools:    make(map[string]*Tool),
				Children: make(map[string]*ToolDir),
			}
			if immediatelyConfigToolDir {
				ConfigToolDir(newToolDir)
			}
			ToolDirMap[tempPathBuilder.String()] = newToolDir
			lastToolDir.Children[s] = newToolDir
		} else {
			newToolDir = tooldir
		}

		lastToolDir = newToolDir
	}

	// 设置tool的ArgSuggests
	argSuggests := make([]prompt.Suggest, 0, len(tool.Args))
	for _, arg := range tool.Args {
		argSuggests = append(argSuggests, prompt.Suggest{Text: arg.Name, Description: arg.Descriptions})
	}
	tool.ArgSuggests = argSuggests
}

func ConfigToolDirs(tooldirs map[string]*ToolDir) {
	// 遍历配置tooldir
	for _, tooldir := range tooldirs {
		ConfigToolDir(tooldir)
	}
}

func ConfigToolDir(tooldir *ToolDir) {
	// 设置tooldir的PathSuggests
	size := len(tooldir.Children)
	if size > 0 {
		suggests := make([]prompt.Suggest, 0, size)
		for name := range tooldir.Children {
			suggests = append(suggests, prompt.Suggest{Text: name})
		}
		tooldir.PathSuggests = suggests
	}

	// 设置tooldir的ToolSuggests
	size = len(tooldir.Tools)
	if size > 0 {
		suggests := make([]prompt.Suggest, 0, size)
		for name, tool := range tooldir.Tools {
			// 对windows做处理
			if utils.IsWindows {
				commandArgs := strings.SplitN(tool.Command, " ", 2)
				binary := strings.TrimSpace(commandArgs[0])
				if !strings.HasSuffix(binary, ".exe") {
					binary += ".exe"
				}

				tooldir.Tools[name].Command = binary + " " + commandArgs[1]
			}
			suggests = append(suggests, prompt.Suggest{Text: name, Description: tool.Descriptions})
		}
		tooldir.ToolSuggests = suggests
	}
}
