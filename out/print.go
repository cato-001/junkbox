package out

import (
	"fmt"
	"os"
)

func Eprintln(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}
