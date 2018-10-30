package main

import "fmt"

func split(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}
	mid := l >> 1
	left := split(a[:mid])
	right := split(a[mid:])
	return merge(left, right)
}

func merge(a []int, b []int) []int {
	temp := make([]int, 0, len(a)+len(b))
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			temp = append(temp, a[i])
			i++
		} else {
			temp = append(temp, b[j])
			j++
		}
	}
	if i < len(a) {
		temp = append(temp, a[i:]...)
	} else if j < len(b) {
		temp = append(temp, b[j:]...)
	}
	return temp
}

func main() {
	var b = []int{27, 38, 12, 39, 27, 16}
	newb := split(b)
	fmt.Println(newb)
}
