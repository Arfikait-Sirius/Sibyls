package sibyls

import (
	"fmt"
)

type emily struct{}

var Emily emily

func (e emily) FnPrintLine(message string) {
	fmt.Println(message)
}
