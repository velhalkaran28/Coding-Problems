package main

import (
	"fmt"
)

func main() {
	test := "This Is A Test To Upper"
	fmt.Println(toUpper(test))
}

func toUpper(str string) string {
	var ret_str = ""
	for _, chr := range str {
		if chr >= 97 && chr <= 122 {
			chr -= 32
		}
		ret_str += fmt.Sprintf("%c", chr)
	}
	return ret_str
}
