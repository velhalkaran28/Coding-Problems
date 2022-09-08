package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6}
	fmt.Println(BinarySearch(arr, 6))
	fmt.Println(BinarySearchRecursive(arr, 0, len(arr)-1, 6))
}

func BinarySearch(arr []int, key int) int {
	l, h := 0, len(arr)-1
	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == key {
			return key
		} else if key < arr[mid] {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func BinarySearchRecursive(arr []int, l, h, key int) int {
	if l <= h {
		mid := (l + h) / 2
		if arr[mid] == key {
			return key
		} else if key < arr[mid] {
			return BinarySearchRecursive(arr, l, mid-1, key)
		} else {
			return BinarySearchRecursive(arr, mid+1, h, key)
		}
	}
	return -1
}
