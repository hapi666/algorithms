package main

import "fmt"

func selectSort(a []int) {
	l := len(a)
	if l <= 1 {
		return
	}
	for i := 0; i < l; i++ {
		temp := i
		for j := i + 1; j < l; j++ {
			if a[temp] > a[j] {
				temp = j
			}
		}
		if temp != i {
			a[i], a[temp] = a[temp], a[i]
		}
	}
}

func main() {
	var b = []int{27, 38, 12, 39, 27, 16}
	selectSort(b)
	fmt.Println(b)
}
