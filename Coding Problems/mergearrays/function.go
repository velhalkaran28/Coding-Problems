package main

import "fmt"

func main() {
	arr1 := []int{3, 8, 16, 20, 25}
	arr2 := []int{4, 10, 12, 22, 23}
	arr3 := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {

		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}
	for ; i < len(arr1); i++ {
		arr3 = append(arr3, arr1[i])
	}
	for ; j < len(arr2); j++ {
		arr3 = append(arr3, arr2[j])
	}

	fmt.Println(arr3)
}
