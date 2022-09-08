package main

import "fmt"

func main() {
	str := "madama"
	flag := true
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			fmt.Println("Not palindrome")
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("Palindrome")
	}
}
