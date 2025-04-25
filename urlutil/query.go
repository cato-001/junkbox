package urlutil

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type (
	Query map[string]any
)

func (q Query) String() string {
	var builder strings.Builder
	first := true
	for key, value := range q {
		if !first {
			builder.WriteRune('&')
		}
		builder.WriteString(url.QueryEscape(key))
		builder.WriteRune('=')
		writeQuery(&builder, value)
		first = false
	}
	return builder.String()
}

func writeQuery(builder *strings.Builder, value any) {
	switch v := value.(type) {
	case string:
		writeQueryEscaped(builder, v)
	case rune:
		writeQueryEscaped(builder, string(v))
	case []byte:
		writeQueryEscaped(builder, string(v))
	case int:
		writeQueryEscaped(builder, strconv.Itoa(v))
	case []any:
		for index, inner := range v {
			if index != 0 {
				builder.WriteRune(',')
			}
			writeQuery(builder, inner)
		}
	default:
		panic(fmt.Sprintf("type not implemented for query value: %T", value))
	}
}

func writeQueryEscaped(builder *strings.Builder, value string) {
	builder.WriteString(url.QueryEscape(value))
}
