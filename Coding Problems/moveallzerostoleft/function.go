package main

import "fmt"

func main() {
	arr := []int{1, 0, 7, 0, 3}
	lastIndex := len(arr) - 1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != 0 {
			arr[lastIndex], arr[i] = arr[i], arr[lastIndex]
			lastIndex--
		}
	}
	fmt.Println(arr)
}
