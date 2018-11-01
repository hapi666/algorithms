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
	for index := (size - 1) / 2; index >= 0; index-- {
		// 对所有父节点，从最后一个父节点到根节点，依次下沉排序
		down(a, index, size)
	}
	// 上浮
	// for index := 0; index < size; index++ {
	// 	up(a, index)
	// }
	for size > 1 {
		a[0], a[size-1] = a[size-1], a[0]
		size--
		down(a, 0, size)
		//up(a, size-1)
	}
}

//上浮
func up(a []int, index int) {
	//从根节点到最后的叶子结点，让他们能上浮的都尽量上浮。
	//遍历每个节点，看它是否大于它的父节点，大于就与父节点进行交换
	for a[index] > a[(index-1)/2] { //这个地方不能用右移一位代替除以2，因为-1右移一位与-1除以2不一样
		a[index], a[(index-1)/2] = a[(index-1)/2], a[index]
		index = (index - 1) / 2
	}
}

//下沉
func down(a []int, index, size int) {
	// 从最后一个父节点（至少存在一个字节点）开始，向下沉
	// 每次都保证此时节点的值大于其支上所有节点的值，就是除了在保证此节点比它左右节点的值大之外，
	// 还要保证它左右节点的 子节点 比 对应的左右节点 小
	// 虽然是在往堆顶遍历，但是每遍历到一个地方，它的分支都要符合大顶堆的要求
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
