package main

import "fmt"

//FastSort is a fast way to sort slice.
func FastSort(a []int) []int {
	if len(a) <= 1 {
		temp := make([]int, len(a))
		copy(temp, a)
		return temp
	}
	pivot := a[0]
	low := make([]int, 0, len(a))
	high := make([]int, 0, len(a))
	for index := 1; index < len(a); index++ {
		if a[index] < pivot {
			low = append(low, a[index])
		} else {
			high = append(high, a[index])
		}
	}
	low, high = FastSort(low), FastSort(high)
	return append(append(low, pivot), high...)
}

func main() {
	//testing data.
	var b = []int{27, 38, 12, 39, 27, 16}
	fmt.Println(FastSort(b))
}
