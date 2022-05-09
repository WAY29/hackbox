package colorprint

import (
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
)

var (
	StandardOutput = os.Stdout

	au aurora.Aurora

	SUCCESS_FLAG string
	ERROR_FLAG   string
	WARNING_FLAG string
)

func InitColorPrint(noColor bool) {
	au = aurora.NewAurora(!noColor)
	SUCCESS_FLAG = fmt.Sprintf("[%s] ", au.Green("+"))
	ERROR_FLAG = fmt.Sprintf("[%s] ", au.Red("-"))
	WARNING_FLAG = fmt.Sprintf("[%s] ", au.Yellow("!"))
}

func Success(fomrat string, args ...interface{}) {
	fmt.Fprintf(StandardOutput, SUCCESS_FLAG+fomrat+"\n", args...)
}

func Error(fomrat string, args ...interface{}) {
	fmt.Fprintf(StandardOutput, ERROR_FLAG+fomrat+"\n", args...)
}

func Warning(fomrat string, args ...interface{}) {
	fmt.Fprintf(StandardOutput, WARNING_FLAG+fomrat+"\n", args...)
}

func Text(format string, args ...interface{}) {
	fmt.Fprintf(StandardOutput, format+"\n", args...)
}

func SetMsg(name, value string, global bool) {
	msg := fmt.Sprintf("%s => %s", au.Cyan(name), value)
	if global {
		msg = fmt.Sprintf("[%s] %s", au.Yellow("global"), msg)
	}
	Success(msg)
}

func ArgumentMsg(name, argType, value string) {
	if len(value) == 0 {
		value = "nil"
	} else {
		au.Bold(value)
	}
	fmt.Fprintf(StandardOutput, "  - %s [%s]: %s\n", au.Cyan(name), au.Yellow(argType), value)
}

func OutputSoonMsg(name string) {
	Success("Outout => %s soon...\n", au.Cyan(name))
}

func OutputMsg(name string) {
	Success("Outout => %s\n", au.Cyan(name))
}

func formatName(name string) string {
	if strings.Contains(name, "$") {
		return name[1:]
	}
	return name
}

func SetOutputMsg(name, value string) {
	name = formatName(name)
	Success("Outout %s => %s\n", au.Cyan("$"+name), value)
}

func ExportMsg(name, path string) {
	name = formatName(name)
	Success("Exrpot %s => %s\n", au.Cyan("$"+name), path)
}
