package out

import (
	"encoding/json"
	"fmt"
)

func Jsonln(data any) {
	content, err := json.Marshal(&data)
	if err != nil {
		Ejsonln(err)
		return
	}
	fmt.Println(string(content))
}

func Ejsonln(err error) {
	type out struct {
		Error string `json:"error"`
	}
	content, err := json.Marshal(&out{Error: err.Error()})
	if err != nil {
		Eprintln("ERROR:", err)
		return
	}
	Eprintln(string(content))
}
