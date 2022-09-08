package main

import "fmt"

func main() {
	str1 := "decimal"
	str2 := "medical"
	flag := true

	if len(str1) != len(str2) {
		fmt.Println("Not anagram")
		return
	}

	var mp = make(map[string]int)
	for _, s := range str1 {
		j := mp[string(s)]
		if j == 0 {
			mp[string(s)] = 1
		} else {
			mp[string(s)] = j + 1
		}
	}

	for _, s := range str2 {
		j := mp[string(s)]
		if j == 0 {
			mp[string(s)] = 1
		} else {
			mp[string(s)] = j + 1
		}
	}

	for _, value := range mp {
		if value%2 != 0 {
			flag = false
			break
		}
	}
	if !flag {
		fmt.Println("Not Anagram")
	} else {
		fmt.Println("Anagram")
	}
}
