package strutil

import "strings"

func Split2(input string, sep string) (string, string) {
	parts := strings.SplitN(input, sep, 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
