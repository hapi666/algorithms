package main

import (
	"fmt"
)

type node struct {
	key    int
	left   *node
	right  *node
	parent *node
}

type Tree struct {
	root *node
}

func (t *Tree) Insert(k int) {
	if t.root == nil {
		t.root = &node{
			key:    k,
			left:   nil,
			right:  nil,
			parent: nil,
		}
	} else {
		t.root.insert(k)
	}
}

func (t *node) insert(k int) {
	if t == nil {
		return
	}
	temp := t
	if k > temp.key {
		if temp.right == nil {
			temp.right = &node{
				key:    k,
				left:   nil,
				right:  nil,
				parent: temp,
			}
			return
		}
		temp.right.insert(k)
	}
	if k < temp.key {
		if temp.left == nil {
			temp.left = &node{
				key:    k,
				left:   nil,
				right:  nil,
				parent: temp,
			}
		}
		temp.left.insert(k)
	}
}

func (t *Tree) Search(k int) *node {
	if t.root == nil {
		return nil
	}
	temp := t.root
	for temp != nil {
		if k == temp.key {
			return temp
		} else if k > temp.key {
			temp = temp.right
		} else if k < temp.key {
			temp = temp.left
		}
	}
	return nil
}

func Min(n *node) *node {
	if n == nil {
		return nil
	}
	temp := n
	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

//就按中序方式找后继结点
func successor(n *node) *node {
	if n.right != nil {
		return Min(n.right)
	}
	temp := n.parent
	for temp != nil && temp == temp.right {
		temp = temp.parent
	}
	return temp
}

func (t *Tree) Delete(k int) {
	if t.root == nil {
		return
	}
	n := t.Search(k)
	var temp *node
	if n.left == nil || n.right == nil {
		temp = n
	} else { //孩子都不为空
		temp = successor(n)
	}
	var x *node
	if temp.left != nil {
		x = temp.left
	} else {
		x = temp.right
	}
	if x != nil { //说明temp有孩子，说明它不是一个叶子结点
		x.parent = temp.parent
	}
	if temp.parent == nil { //说明temp（后继）是根结点 或者 此时你要删除的结点就是根结点
		t.root = x
	} else {
		if temp == temp.parent.left {
			temp.parent.left = x
		} else {
			temp.parent.right = x
		}
		if temp != n { //是后继那种情况
			n.key = temp.key
		}
	}
}

func (t *Tree) Display() {
	if t.root == nil {
		return
	} else {
		t.root.display()
	}
}

//中序
func (n *node) display() {
	if n != nil {
		n.left.display()
		fmt.Println(n.key)
		n.right.display()
	}
}

func main() {
	t := new(Tree)
	t.Insert(7)
	t.Insert(1)
	t.Insert(3)
	t.Insert(9)
	t.Insert(2)
	t.Insert(6)
	t.Display()
	if t.Search(7) != nil {
		fmt.Println("search result:")
		fmt.Println(t.Search(7).key)
	}
	t.Delete(6)
	fmt.Println("delete result:")
	t.Display()
}
