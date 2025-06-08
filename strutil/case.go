package strutil

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CamelCase(input string) string {
	input = PascalCase(input)
	return strings.ToLower(string(input[0])) + input[1:]
}

func PascalCase(input string) string {
	dashReplacer := strings.NewReplacer("-", " ", "_", " ")
	caser := cases.Title(language.German, cases.NoLower)
	spaceReplacer := strings.NewReplacer("\t", "", " ", "")
	return spaceReplacer.Replace(caser.String(dashReplacer.Replace(input)))
}

func UpperSnakeCase(input string) string {
	caser := cases.Upper(language.German)
	return caser.String(SnakeCase(input))
}

func SnakeCase(input string) string {
	return delimited(input, "_")
}

func KebabCase(input string) string {
	return delimited(input, "-")
}

func delimited(input, delimiter string) string {
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyzöäü"
	uppercaseLetters := strings.ToUpper(lowercaseLetters)

	replacements := make([]string, len(lowercaseLetters) * 2)
	for index, lowercaseRune := range lowercaseLetters {
		lowercaseLetter := string(lowercaseRune)
		uppercaseLetter := string(uppercaseLetters[index])
		replacements[index * 2] = uppercaseLetter
		replacements[index * 2 + 1] = delimiter + lowercaseLetter
	}

	uppercaseReplacer := strings.NewReplacer(replacements...)
	spaceReplacer := strings.NewReplacer("-", delimiter, "_", delimiter, "\t", delimiter, " ", delimiter)

	input = uppercaseReplacer.Replace(input)
	input = spaceReplacer.Replace(input)
	return strings.Trim(input, "-_\t ")
}
