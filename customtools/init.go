package customtools

import (
	"github.com/WAY29/hackbox/customtools/test"
	_ "github.com/WAY29/hackbox/customtools/test"
	"github.com/WAY29/hackbox/internal/tools"
)

func InitCustomTools() {
	tools.NewTool(
		"other.test", // tools path
		"(builtin) just a test, print all arguments", // tool descriptions
		"str",    // tools args expression
		test.Run, // function that will be run, Function signature must be func(map[string]string)
		// add arguments for tool
		*tools.NewArguments("str", "string", "just a string"),
	)
}
