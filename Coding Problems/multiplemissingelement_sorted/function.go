package main

import "fmt"

func main() {
	arr := []int{6, 7, 8, 9, 11, 12, 15, 16, 17, 18, 19}
	diff := arr[0] - 0
	for i := 0; i < len(arr); i++ {
		if arr[i]-i != diff {
			for diff < arr[i]-i {
				fmt.Printf("Missing element is %d\n", i+diff)
				diff++
			}
		}
	}
}
