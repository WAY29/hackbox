package main

import (
	"flag"

	"github.com/WAY29/hackbox/customtools"
	_ "github.com/WAY29/hackbox/customtools" // 为了触发init注入自定义工具
	"github.com/WAY29/hackbox/internal/cmd"
	_ "github.com/WAY29/hackbox/internal/cmd/root"  // 为了触发init注入命令
	_ "github.com/WAY29/hackbox/internal/cmd/tools" // 为了触发init注入命令

	"github.com/WAY29/hackbox/internal/colorprint"
	. "github.com/WAY29/hackbox/internal/colorprint"
	"github.com/WAY29/hackbox/internal/prompts"
	"github.com/WAY29/hackbox/internal/prompts/completer"
	"github.com/WAY29/hackbox/internal/prompts/executor"
	"github.com/WAY29/hackbox/internal/tools"
	"github.com/WAY29/hackbox/utils"

	"github.com/c-bata/go-prompt"
)

const (
	__version__ = "1.1.0"
)

var (
	noColor        bool
	quiet          bool
	customToolPath string
)

func init() {
	flag.BoolVar(&noColor, "nc", false, "Print without color")
	flag.BoolVar(&quiet, "q", false, "Run hackbox without banner")
	flag.StringVar(&customToolPath, "p", "", "Custom tool path, default will load from ./tools.toml -> $HOME/.config/hackbox/tools.yaml")
}

func banner() {
	Text(`
    _   _   ___  _____  _   ________  _______   __
   | | | | / _ \/  __ \| | / /| ___ \|  _  \ \ / /
   | |_| |/ /_\ \ /  \/| |/ / | |_/ /| | | |\ V / 
   |  _  ||  _  | |    |    \ | ___ \| | | |/   \ 
   | | | || | | | \__/\| |\  \| |_/ /\ \_/ / /^\ \
   \_| |_/\_| |_/\____/\_| \_/\____/  \___/\/   \/

   `)
}

func main() {
	// 解析标志
	flag.Parse()
	// 输出banner
	if !quiet {
		banner()
	}
	// 初始化colorprint
	colorprint.InitColorPrint(noColor)
	// 检查是否以root权限启动并输出警告
	utils.CheckRoot(quiet)
	// 初始化tools.yaml
	if err := tools.LoadTools(tools.GetToolPath(customToolPath), quiet); err != nil {
		Error(err.Error())
		return
	}
	// 初始化自定义工具
	customtools.InitCustomTools()
	// 初始化prompt
	prompts.InitPrompt("/")
	// 初始化cmd
	cmd.InitCmd()

	// 清理
	defer func() {
		// 清理临时文件
		for _, name := range utils.TempFileNames {
			utils.RemoveFile(name)
		}
	}()

	p := prompt.New(
		executor.Executor,
		completer.Completer,
		prompt.OptionPrefix("hackbox> "),
		prompt.OptionLivePrefix(prompts.LivePrefix),
	)

	p.Run()
}
