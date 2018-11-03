package main

import "fmt"

func insertSearch(a []int, find int) int {
	low := 0
	high := len(a) - 1
	for a[low] != a[high] && find >= a[low] && find <= a[high] {
		mid := low + (high-low)*(find-a[low])/(a[high]-a[low])
		if find > a[mid] {
			low = mid + 1
		} else if find < a[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	if a[low] == find {
		return low
	}
	return -1
}

func main() {
	var b = []int{12, 16, 27, 27, 38, 39}
	location := insertSearch(b, 38)
	fmt.Println(location)
}
