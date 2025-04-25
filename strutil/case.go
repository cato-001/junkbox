package strutil

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PascalCase(input string) string {
	dashReplacer := strings.NewReplacer("-", " ", "_", " ")
	caser := cases.Title(language.German, cases.NoLower)
	spaceReplacer := strings.NewReplacer("\t", "", " ", "")
	return spaceReplacer.Replace(caser.String(dashReplacer.Replace(input)))
}
