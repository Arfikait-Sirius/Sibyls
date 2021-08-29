package sibyls

import (
	"strings"
)

type louise struct{}

var Louise louise

func (l louise) FnUpperAll(base string) string {
	return strings.ToUpper(base)
}
