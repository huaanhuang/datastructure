package list

import (
	"errors"
	"fmt"
)

// Node 双链表结构定义
type Node[T any] struct {
	Val  T
	Next *Node[T]
	Last *Node[T]
}

func NewLinkList[T any](maxSizeParam ...int) IList[T] {
	maxSize := -1
	if len(maxSizeParam) > 0 {
		maxSize = maxSizeParam[0]
	}
	return &linkList[T]{maxSize: maxSize}
}

type linkList[T any] struct {
	head    *Node[T] // 头节点
	tail    *Node[T] // 尾节点
	size    int      // 当前大小
	maxSize int      // 最大容量
}

// Size 长度
func (obj *linkList[T]) Size() uint {
	return uint(obj.size)
}

// IsEmpty 是否为空
func (obj *linkList[T]) IsEmpty() bool {
	return obj.size == 0
}

// IsFull 是否满
func (obj *linkList[T]) IsFull() bool {
	if obj.maxSize < 0 {
		return false
	}
	return obj.size == obj.maxSize
}

// HeadPush 头部插入
func (obj *linkList[T]) HeadPush(val T) error {
	if obj.IsFull() {
		return fmt.Errorf("exceeding the maximum length: %d", obj.maxSize)
	}
	obj.size++
	node := &Node[T]{Val: val}
	if obj.head == nil {
		obj.head = node
		obj.tail = node
	} else {
		node.Next = obj.head
		obj.head.Last = node
		obj.head = node
	}
	return nil
}

// HeadPop 头部弹出
func (obj *linkList[T]) HeadPop() (T, error) {
	if obj.IsEmpty() {
		return *new(T), errors.New("container is empty")
	}
	obj.size--
	node := obj.head
	obj.head = obj.head.Next
	if obj.head == nil {
		obj.tail = nil
	} else {
		obj.head.Last = nil
	}
	return node.Val, nil
}

// TailPush 尾部插入
func (obj *linkList[T]) TailPush(val T) error {
	if obj.IsFull() {
		return fmt.Errorf("exceeding the maximum length: %d", obj.maxSize)
	}
	obj.size++
	node := &Node[T]{Val: val}
	if obj.tail == nil {
		obj.head = node
		obj.tail = node
	} else {
		obj.tail.Next = node
		node.Last = obj.tail
		obj.tail = node
	}
	return nil
}

// TailPop 尾部弹出
func (obj *linkList[T]) TailPop() (T, error) {
	if obj.IsEmpty() {
		return *new(T), errors.New("container is empty")
	}
	obj.size--
	node := obj.tail
	obj.tail = obj.tail.Last
	if obj.tail == nil {
		obj.head = nil
	} else {
		obj.tail.Next = nil
	}
	return node.Val, nil
}

// HeadPeek 获取头部元素
func (obj *linkList[T]) HeadPeek() (T, error) {
	if obj.IsEmpty() {
		return *new(T), errors.New("container is empty")
	}
	return obj.head.Val, nil
}

// TailPeek 获取尾部元素
func (obj *linkList[T]) TailPeek() (T, error) {
	if obj.IsEmpty() {
		return *new(T), errors.New("container is empty")
	}
	return obj.tail.Val, nil
}

// Iteration 迭代容器元素
func (obj *linkList[T]) Iteration(iterationFunc IterationFunc[T], directionParam ...IterationDirection) error {
	direction := StartToFinish
	if len(directionParam) > 0 {
		direction = directionParam[0]
	}
	if direction == StartToFinish {
		current := obj.head
		for current != nil {
			err := iterationFunc(current.Val)
			if err != nil {
				return err
			}
			current = current.Next
		}
	} else {
		current := obj.tail
		for current != nil {
			err := iterationFunc(current.Val)
			if err != nil {
				return err
			}
			current = current.Last
		}
	}
	return nil
}

// ToArray 转换为数组
func (obj *linkList[T]) ToArray() []T {
	var (
		arr     = make([]T, obj.size)
		index   = 0
		current = obj.head
	)
	for current != nil {
		arr[index] = current.Val
		index++
		current = current.Next
	}
	return arr
}

// Push 入栈
func (obj *linkList[T]) Push(item T) error {
	return obj.HeadPush(item)
}

// Pop 出栈
func (obj *linkList[T]) Pop() (item T, err error) {
	return obj.HeadPop()
}

// Peek 返回栈顶元素
func (obj *linkList[T]) Peek() (item T, err error) {
	return obj.HeadPeek()
}

// Qsize 队列长度
func (obj *linkList[T]) Qsize() uint {
	return uint(obj.size)
}

// Put 入队
func (obj *linkList[T]) Put(item T) error {
	return obj.HeadPush(item)
}

// Get 出队
func (obj *linkList[T]) Get() (T, error) {
	return obj.TailPop()
}
