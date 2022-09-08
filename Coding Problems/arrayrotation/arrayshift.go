package main

import "fmt"

func main() {
	arr := []int{8, 3, 9, 15, 10, 7, 2, 12, 4}
	d := 2
	var element int
	for d > 0 {
		element, arr = arr[0], arr[1:]
		fmt.Println(element, " getting rotated")
		arr = append(arr, element)
		d--
	}
	fmt.Println("After rotate: ", arr)
}
