package main

import "fmt"

func main() {
	arr1 := []int{32, 57, 35, 22}
	fmt.Println(insert(arr1, 99, 2))
}

func insert(arr []int, element, pos int) []int {
	return append(arr[:pos], append([]int{element}, arr[pos:]...)...)
}
