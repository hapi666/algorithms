package main

import (
	"fmt"
	"log"
)

type stack struct {
	data   int
	next   *stack
	length int
}

func newStack(l int) *stack {
	s := new(stack)
	s.length = l
	s.next = nil
	return s
}

func (s *stack) isEmpty() bool {
	return s.length == 0
}

func (s *stack) Push(d int) {
	temp := new(stack)
	temp.data = d
	temp.next = s.next
	s.next = temp
}

func (s *stack) Pop() int {
	if s.isEmpty() {
		log.Println("空栈")
		return 0
	}
	p := s.next.data
	s.next = s.next.next
	return p
}

func main() {
	S := newStack(10)
	for index := 0; index < 10; index++ {
		S.Push(index)
	}
	for index := 0; index < 10; index++ {
		fmt.Println(S.Pop())
	}
}
