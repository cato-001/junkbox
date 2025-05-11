package nom

import (
	"fmt"
	"strings"
)

func Rune(value rune) RuneParser {
	return func(input string) (string, rune, error) {
		after, found := strings.CutPrefix(input, string(input))
		if !found {
			return input, value, fmt.Errorf("could not find rune: %v", value)
		}
		return after, value, nil
	}
}

func Str(value string) StrParser {
	return func(input string) (string, string, error) {
		after, found := strings.CutPrefix(input, value)
		if !found {
			return input, value, fmt.Errorf("could not find str: %s", value)
		}
		return after, value, nil
	}
}
