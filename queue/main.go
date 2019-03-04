package main

import (
	"fmt"
	"log"
)

type node struct {
	data int
	next *node
}

type Queue struct {
	front *node
	rear  *node
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.front = new(node)
	queue.rear = new(node)
	queue.front = queue.rear
	return queue
}

// 入队列
func (q *Queue) EnQueue(newData int) {
	temp := new(node)
	temp.data = newData
	temp.next = nil
	q.rear.next = temp
	q.rear = temp
}

// 出队列
func (q *Queue) DeQueue() int {
	if q.front.next == q.rear.next {
		log.Println("空栈无法弹出")
		return -1
	}
	val := q.front.next.data
	q.front.next = q.front.next.next
	return val
}

func main() {
	queue := NewQueue()
	for index := 1; index <= 10; index++ {
		queue.EnQueue(index)
	}
	for index := 0; index < 10; index++ {
		fmt.Println(queue.DeQueue())
	}
}
