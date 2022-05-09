package prompts

import "fmt"

func LivePrefix() (string, bool) {
	if len(Prompt.UseToolName) > 0 {
		return fmt.Sprintf("hackbox %s [%s]> ", Prompt.Path, Prompt.UseToolName), true
	}
	return fmt.Sprintf("hackbox %s> ", Prompt.Path), true
}
