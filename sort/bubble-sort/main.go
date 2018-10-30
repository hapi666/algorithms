package main

import "fmt"

func bubbleSort(a []int) {
	l := len(a)
	if l <= 1 {
		return
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if a[j+1] < a[j] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
}

func main() {
	var b = []int{27, 38, 12, 39, 27, 16}
	bubbleSort(b)
	fmt.Println(b)
}
