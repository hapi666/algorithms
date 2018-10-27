package main

import "fmt"

func insertSort(a []int) []int {
	l := len(a)
	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
	temp := make([]int, l)
	copy(temp, a)
	return temp
}

func main() {
	var a = []int{27, 38, 12, 39, 27, 16}
	newa := insertSort(a)
	fmt.Println(newa)
}
