package util

import (
	"fmt"
	"strings"
)

func Assert(cond bool, msg ...any) {
	if !cond {
		if len(msg) > 0 {
			panic(strings.TrimSpace(fmt.Sprint(msg...)))
		} else {
			panic("Assertion failed")
		}
	}
}
