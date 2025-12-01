package printer

import "fmt"

var enabled bool = true

func Enable() {
	enabled = true
}

func Disable() {
	enabled = false
}

func Ln(a ...any) {
	if enabled {
		fmt.Println(a...)
	}
}

func F(format string, args ...interface{}) {
	if enabled {
		fmt.Printf(format, args...)
	}
}
