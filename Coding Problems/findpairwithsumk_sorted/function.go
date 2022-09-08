package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5, 6, 8, 9, 10, 12, 14}
	sum := 10

	i, j := 0, len(arr)-1
	flag := false
	for i < j {
		if arr[i]+arr[j] == sum {
			fmt.Printf("Pair (%d, %d)\n", arr[i], arr[j])
			i++
			j--
			flag = true
		} else if arr[i]+arr[j] < sum {
			i++
		} else {
			j--
		}
	}
	if !flag {
		fmt.Println("No pair found")
	}
}
