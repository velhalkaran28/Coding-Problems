package main

import "fmt"

var stack = make([]string, 0)

func main() {
	str := "(a+b]-c*(d/2)"
	var mp = map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	for _, value := range str {
		if string(value) == "{" || string(value) == "[" || string(value) == "(" {
			stack = append(stack, string(value))
		} else if string(value) == "}" || string(value) == "]" || string(value) == ")" {
			if len(stack) == 0 {
				fmt.Println("Invalid expression")
				return
			}
			poppedElement := stack[len(stack)-1]
			if poppedElement == mp[string(value)] {
				stack = stack[:len(stack)-1]
			} else {
				fmt.Println("Invalid expression")
				return
			}
		}
	}
	fmt.Println("Valid expression")
}
