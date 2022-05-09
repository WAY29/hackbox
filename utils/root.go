package utils

import (
	"os"
	"runtime"

	. "github.com/WAY29/hackbox/internal/colorprint"
)

func CheckRoot(quiet bool) {
	if runtime.GOOS == "window" || quiet {
		return
	}

	if os.Geteuid() != 0 {
		Warning("You run hackbox without root privilege. Some program may not work.")
	}
}
