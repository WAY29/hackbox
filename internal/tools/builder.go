package tools

type ToolBuilder struct{}

type ArgumentsFunc func(*Tool)

func NewTool(toolPath, descriptions string, ArgsExpression string, builtinFunc BuiltinFunc, args ...Arg) {
	ConfigTool(toolPath, &Tool{
		Descriptions:   descriptions,   // tool description
		Args:           args,           // add arguments for tool, Only three fields are required: Name, Type, Descriptions
		ArgsExpression: ArgsExpression, // run failed if expression is false
		BuiltinFunc:    builtinFunc,    // builtin function, Function signature must be func(map[string]string)
	}, true)

}

func NewArguments(name, argType, descriptions string) *Arg {
	return &Arg{
		Name:         name,
		Type:         argType,
		Descriptions: descriptions,
	}
}
