package main

import (
	"fmt"
)

func BinarySearch(Arry []int, number int) int {
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
	count := 5
	i := 7
	arry := make([]int, 5)
	for index := 0; index < count; index++ {
		arry[index] = i
		i++
	}
	fmt.Println(BinarySearch(arry, 7))
}

