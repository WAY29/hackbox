package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

var (
	IsWindows   = false
	shellPath   string
	shellRunArg string
)

func init() {
	if runtime.GOOS == "windows" {
		IsWindows = true
		shellPath, _ = exec.LookPath("cmd.exe")
		shellRunArg = "/c"
	} else {
		shellPath, _ = exec.LookPath("sh")
		shellRunArg = "-c"
	}
}

func ExecCommand(command string) *exec.Cmd {
	args := []string{shellRunArg, command}
	cmd := exec.Command(shellPath, args...)

	return cmd
}

func main() {
	r := ExecCommand("echo $a")
	r.Env = []string{"a=b\nqwe"}
	// fmt.Printf("%#v\n", )

	// if err != nil {
	// 	panic(err)
	// }
	output, err := r.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
}
