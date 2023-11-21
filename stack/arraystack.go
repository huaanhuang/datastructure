package stack

import (
	"errors"
	"fmt"
	"github.com/huaanhuang/datastructure/port"
)

// NewArrayStack 栈(数组实现)
// cap: 栈容量
func NewArrayStack[T any](cap uint) port.IStack[T] {
	return &arrayStack[T]{
		data: make([]T, cap, cap),
		size: cap,
		cur:  -1,
	}
}

type arrayStack[T any] struct {
	data   []T  // 栈容器
	size   uint // 栈容器大小
	length uint // 当前栈长度
	cur    int  // 当前栈顶元素索引
}

// Push 入栈
func (obj *arrayStack[T]) Push(item T) error {
	if obj.length == obj.size {
		return errors.New(fmt.Sprintf("exceeds stack capacity %d", obj.size))
	}
	obj.cur++
	obj.data[obj.cur] = item
	obj.length++
	return nil
}

// Pop 出栈
func (obj *arrayStack[T]) Pop() (item T, err error) {
	if obj.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	ans := obj.data[obj.cur]
	obj.length--
	obj.cur--
	return ans, nil
}

// IsEmpty 栈是否为空
func (obj *arrayStack[T]) IsEmpty() bool {
	return obj.length == 0
}

// Size 栈当前大小
func (obj *arrayStack[T]) Size() uint {
	return obj.length
}

// Peek 返回栈顶元素
func (obj *arrayStack[T]) Peek() (item T, err error) {
	if obj.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	return obj.data[obj.cur], nil
}
