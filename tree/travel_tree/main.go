package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

func createTree(nodes []int) *tree {
	if len(nodes) == 0 {
		return &tree{}
	}
	temp := new(tree)
	temp.val = nodes[0]
	for _, v := range nodes {
		insertNode(v, temp)
	}
	return temp
}

func insertNode(v int, t *tree) {
	if t == nil {
		fmt.Println("111")
		t = &tree{val: v}
	}
	if v > t.val {
		if t.right == nil {
			t.right = &tree{val: v}
		} else {
			insertNode(v, t.right)
		}

	} else if v < t.val {
		if t.left == nil {
			t.left = &tree{val: v}
		} else {
			insertNode(v, t.left)
		}
	}
}

func middleTravel(t *tree) {
	if t == nil {
		return
	}

	middleTravel(t.left)
	fmt.Println(t.val)
	middleTravel(t.right)
}

func beforeTravel(t *tree) {
	if t == nil {
		return
	}

	fmt.Println(t.val)
	beforeTravel(t.left)
	beforeTravel(t.right)
}

func afterTravel(t *tree) {
	if t == nil {
		return
	}

	afterTravel(t.left)
	afterTravel(t.right)
	fmt.Println(t.val)
}

func main() {
	test := []int{7, 4, 9, 1, 8, 2, 3}
	t := createTree(test)
	//fmt.Println(t.val)
	fmt.Println("中序遍历开始～")
	middleTravel(t)
	fmt.Println("先序遍历开始～")
	beforeTravel(t)
	fmt.Println("后序遍历开始～")
	afterTravel(t)
	fmt.Println()
	insertNode(0, t)
	middleTravel(t)
}
