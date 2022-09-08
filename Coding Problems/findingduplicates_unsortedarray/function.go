package main

import "fmt"

func main() {
	arr1 := []int{8, 3, 6, 4, 6, 5, 6, 8, 2, 7}
	arr2 := []int{}
	mp := make(map[int]bool)
	for _, e := range arr1 {
		if !mp[e] {
			mp[e] = true
			arr2 = append(arr2, e)
		}
	}
	fmt.Println(arr2)
}
