package queue

// IQueue 队列接口
type IQueue[T any] interface {
	Qsize() uint
	IsEmpty() bool
	IsFull() bool
	Put(item T) error
	Get() (T, error)
}
