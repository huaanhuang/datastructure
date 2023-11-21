package queue

import (
	"errors"
	"time"
)

func NewArrayQueue[T any](size uint) IQueue[T] {
	qsize := size + 1
	return &arrayQueue[T]{
		items: make([]T, qsize, qsize),
		size:  qsize,
		head:  0,
		tail:  0,
	}
}

type arrayQueue[T any] struct {
	items []T  // 队列元素
	size  uint // 队列长度
	head  uint // 头指针
	tail  uint // 尾指针
}

// Qsize 队列当前长度
func (q *arrayQueue[T]) Qsize() uint {
	return q.tail - q.head
}

// IsEmpty 队列是否为空
func (q *arrayQueue[T]) IsEmpty() bool {
	return q.head == q.tail
}

// IsFull 队列是否已满
func (q *arrayQueue[T]) IsFull() bool {
	return (q.tail+1)%q.size == q.head
}

// Put 入队
func (q *arrayQueue[T]) Put(item T, timeout time.Duration) error {
	// 判断队列是否已满
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.size // 移到下一个节点
	return nil
}

// Get 出队
func (q *arrayQueue[T]) Get(timeout time.Duration) (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("queue is empty")
	}
	val := q.items[q.head]
	q.head = (q.head + 1) % q.size // 移动到下一个节点
	return val, nil
}
