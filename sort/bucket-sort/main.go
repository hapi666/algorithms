package main

import "fmt"

func bucketSort(a []int, m int) {
	if len(a) <= 1 {
		return
	}
	bucket := make([]int, m)
	for index := 0; index < len(a); index++ {
		bucket[a[index]]++
	}
	for index := 0; index < m; index++ {
		for bucket[index] >= 1 {
			fmt.Printf("%d ", index)
			bucket[index]--
		}
	}
}

func main() {
	var b = []int{27, 38, 12, 39, 27, 16}
	bucketSort(b, 40)
}
