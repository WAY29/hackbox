package utils

import "fmt"

func StringEllipsis(str string, length int) string {
	if len(str) > length {
		str = str[:length] + "..."
	}

	str = fmt.Sprintf("%#v", str)
	str = str[1 : len(str)-1]

	return str
}
