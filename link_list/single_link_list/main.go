package main

import (
	"fmt"
	"log"
)

type Node struct {
	key    interface{}
	next   *Node
	parent *Node
}

func (node *Node) insertNode(insertKey interface{}) {
	head := node
	if head == nil {
		head = &Node{
			key:    insertKey,
			next:   nil,
			parent: nil,
		}
		return
	}
	newNode := new(Node)
	for head.next != nil {
		head = head.next
	}
	head.next = new(Node)
	newNode.parent = head
	newNode.key = insertKey
	head.next = newNode
}

func (node *Node) searchNode(searchKey interface{}) *Node {
	if node == nil {
		log.Println("链表为空，查找失败")
		return nil
	}
	temp := node
	for temp != nil && temp.key != searchKey {
		temp = temp.next
	}
	if temp == nil {
		return nil
	}
	return temp
}

func (node *Node) deleteNode(deleteKey interface{}) {
	sn := node.searchNode(deleteKey)
	if sn == nil {
		return
	}
	*sn.parent.next = *sn.next
	*sn.next.parent = *sn.parent
}

func (node *Node) Display() {
	temp := node
	for temp.next != nil {
		temp = temp.next
		fmt.Printf("%v -> ", temp.key)
	}
	fmt.Println()
}

func main() {
	n := new(Node)
	n.insertNode(7)
	n.insertNode(9)
	n.insertNode(5)
	n.insertNode(1)
	n.insertNode(4)
	n.insertNode(2)
	n.Display()
	n.deleteNode(1)
	n.Display()
	fmt.Println(n.searchNode(9).key)
}
