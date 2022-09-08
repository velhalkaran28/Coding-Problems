package main

import "fmt"

func main() {
	arr := []int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20, 20}
	lastDuplicate := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] && arr[i] != lastDuplicate {
			fmt.Printf("%d\t", arr[i])
			lastDuplicate = arr[i]
		}
		if arr[i] == arr[i+1] {
			j := i + 1
			for arr[j] == arr[i] {
				j++
			}
			fmt.Printf("%d is found %d times\n", arr[i], j-i)
			i = j - 1
		}
	}
}
