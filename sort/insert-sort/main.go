package main

import "fmt"

func insertSort(a []int) {
	l := len(a)
	if l <= 1 {
		return
	}
	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func main() {
	var b = []int{27, 38, 12, 39, 27, 16}
	insertSort(b)
	fmt.Println(b)
}
