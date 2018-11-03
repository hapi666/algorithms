package main

import (
	"fmt"
)

func binarySearch(Arry []int, number int) int {
	minIndex := 0
	maxIndex := len(Arry) - 1
	for minIndex <= maxIndex {
		midIndex := (minIndex + maxIndex) / 2
		midItem := Arry[midIndex]
		if number == midItem {
			return midIndex
		}
		if number < midItem {
			maxIndex = midIndex - 1
		} else if number > midItem {
			minIndex = midIndex + 1
		}
	}
	return -1
}

func main() {
	var b = []int{12, 16, 27, 27, 38, 39}
	location := binarySearch(b, 38)
	fmt.Println(location)
}
