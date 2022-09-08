package main

import "fmt"

func main() {
	arr := []int{6, 3, 8, 10, 16, 7, 5, 2, 9, 14}
	mp := make(map[int]int)
	sum := 10
	flag := false
	for i := 0; i < len(arr); i++ {
		if mp[sum-arr[i]] == 0 {
			mp[arr[i]] = i
		} else {
			fmt.Printf("Pair = (%d , %d)", arr[mp[sum-arr[i]]], arr[i])
			flag = true
		}
	}
	if !flag {
		fmt.Println("No pair found")
	}
}
