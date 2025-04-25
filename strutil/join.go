package strutil

import "strings"

func JoinLastSep(values []string, sep string, lastSep string) string {
	if len(values) == 0 {
		return ""
	}
	if len(values) == 1 {
		return values[0]
	}
	return strings.Join(values[:len(values)-1], sep) + lastSep + values[len(values)-1]
}
