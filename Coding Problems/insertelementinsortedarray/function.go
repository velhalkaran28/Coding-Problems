package main

import "fmt"

func main() {
	arr := []int{4, 8, 13, 15, 17, 21, 25, 27, 31}
	element := 22
	i := len(arr) - 1
	arr = append(arr, 0)
	for arr[i] > element {
		arr[i+1] = arr[i]
		i--
	}
	arr[i+1] = element
	fmt.Println(arr)
}
