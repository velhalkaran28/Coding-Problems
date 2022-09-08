package main

import "fmt"

func main() {
	num := 0
	sum := 0
	for num > 0 {
		rem := num % 10
		sum += rem
		num = num / 10
	}
	fmt.Println(sum)
}
