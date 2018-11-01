package main

import (
	"fmt"
	"time"
)

func heapSort(a []int) {
	if len(a) <= 1 {
		return
	}
	size := len(a)
	intoBigHeap(a, size-1)
	for size > 1 {
		a[0], a[size-1] = a[size-1], a[0]
		size--
		//down(a, 0, size)
		intoBigHeap(a, size)
	}
}

//上浮
func intoBigHeap(a []int, size int) {
	for index := 0; index < size; index++ {
		// if a[index] > a[(index-1)/2] {
		// 	a[index], a[index*2+1] = a[index*2+1], a[index]
		// }
		//从根节点到最后的叶子结点，让他们能上浮的都尽量上浮。
		//遍历每个节点，看它是否大于它的父节点，大于就与父节点进行交换
		for a[index] > a[(index-1)/2] {
			a[index], a[(index-1)/2] = a[(index-1)/2], a[index]
			index = (index - 1) / 2
		}
		//至此就完成了大顶堆的构建。
	}
}

//下沉
func down(a []int, index, size int) {
	kid := index*2 + 1
	var largest int
	for kid <= size-1 {
		if kid == size-1 {
			largest = kid
			goto label
		}
		if a[kid] > a[kid+1] {
			largest = kid
		} else {
			largest = kid + 1
		}
	label:
		if a[largest] < a[index] {
			break
		}
		a[largest], a[index] = a[index], a[largest]
		index = largest
		kid = index*2 + 1
	}
}

func main() {
	now := time.Now()
	var b = []int{27, 38, 12, 39, 27, 16}
	heapSort(b)
	fmt.Println(b)
	fmt.Printf("%v\n", time.Since(now))
}
