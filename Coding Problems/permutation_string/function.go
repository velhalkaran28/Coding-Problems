package main

import "fmt"

func main() {
	str := "ABC"
	A := make([]int, len(str))
	res := make([]rune, len(str))
	perm([]rune(str), 0, A, res)
	fmt.Println("----------------")
	perm2([]rune(str), 0, len(str)-1)
}

func perm(s []rune, k int, A []int, res []rune) {
	if k == len(s) {
		fmt.Println(string(res))
	} else {
		for i := 0; i != len(s); i++ {
			if A[i] == 0 {
				res[k] = s[i]
				A[i] = 1
				perm(s, k+1, A, res)
				A[i] = 0
			}
		}
	}
}

func perm2(s []rune, l, h int) {
	if l == h {
		fmt.Println(string(s))
	} else {
		for i := l; i <= h; i++ {
			s[l], s[i] = s[i], s[l]
			perm2(s, l+1, h)
			s[l], s[i] = s[i], s[l]
		}
	}
}
