package main

import "fmt"

func main() {
	arr1 := []int{6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 17}
	diff := arr1[0] - 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i]-i != diff {
			fmt.Printf("Missing element is %d", i+diff)
			break
		}
	}
}
