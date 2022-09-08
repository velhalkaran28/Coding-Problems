package main

import "fmt"

func main() {
	str := "hello world"
	rns := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	fmt.Println(string(rns))
	ReverseRecursion(str)
}

func ReverseRecursion(str string) {
	if str == "" {
		return
	}
	ReverseRecursion(str[1:])
	fmt.Print(string(str[0]))
}
