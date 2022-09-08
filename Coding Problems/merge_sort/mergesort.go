package main

import "fmt"

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left := mergeSort(a[:len(a)/2])
	right := mergeSort(a[len(a)/2:])
	return merge(left, right)
}
func merge(a []int, b []int) []int {
	final := make([]int, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
func main() {
	arr := []int{6, 2, 9, 0, 1, 3}
	fmt.Println(mergeSort(arr))
}
