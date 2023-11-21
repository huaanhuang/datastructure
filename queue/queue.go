package queue

import "time"

// IQueue 队列接口
type IQueue[T any] interface {
	Qsize() uint
	IsEmpty() bool
	IsFull() bool
	Put(item T, timeout time.Duration) error
	Get(timeout time.Duration) (T, error)
}
