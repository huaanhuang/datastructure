package stack

import (
	"errors"
	"fmt"
)

func NewLinkStack[T any](maxSizeParam ...int) IStack[T] {
	maxSize := -1
	if len(maxSizeParam) > 0 {
		maxSize = maxSizeParam[0]
	}
	return &linkStack[T]{maxlength: maxSize}
}

type Node[T any] struct {
	Val  T        // 当前节点值
	Next *Node[T] // 下一个节点指针
}

type linkStack[T any] struct {
	head      *Node[T] // 头节点
	length    uint     // 栈大小
	maxlength int      // 栈最大大小
}

// Push 入栈
func (obj *linkStack[T]) Push(item T) error {
	if obj.IsFull() {
		return fmt.Errorf("stack is full with %d element", obj.maxlength)
	}
	node := &Node[T]{Val: item}
	if obj.head == nil {
		obj.head = node
	} else {
		node.Next = obj.head
		obj.head = node
	}
	obj.length++
	return nil
}

// Pop 出栈
func (obj *linkStack[T]) Pop() (item T, err error) {
	if obj.head == nil {
		return *new(T), errors.New("stack is empty")
	}
	item = obj.head.Val
	obj.head = obj.head.Next
	obj.length--
	return item, nil
}

// IsEmpty 栈是否为空
func (obj *linkStack[T]) IsEmpty() bool {
	return obj.length == 0
}

// IsFull 栈是否为满
func (obj *linkStack[T]) IsFull() bool {
	return int(obj.length) == obj.maxlength
}

// Size 栈大小
func (obj *linkStack[T]) Size() uint {
	return obj.length
}

// Peek 返回栈顶元素
func (obj *linkStack[T]) Peek() (item T, err error) {
	if obj.head == nil {
		return *new(T), errors.New("stack is empty")
	}
	return obj.head.Val, nil
}
