package example

import (
	"fmt"
)

var (
	s string
	t = "ttt"
)

const foo = "bar"

func x() {
	fmt.Println(yzzy("x!"))
}

func yzzy(s string) string {
	return s
}
