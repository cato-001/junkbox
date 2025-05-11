package nom

type (
	Parser[T any] func(string) (string, T, error)

	AnyParser Parser[any]
	RuneParser Parser[rune]
	IntParser Parser[int]
	StrParser Parser[string]
)

