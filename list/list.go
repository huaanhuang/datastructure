package list

import (
	"github.com/huaanhuang/datastructure/queue"
	"github.com/huaanhuang/datastructure/stack"
)

// IterationFunc 迭代方法
type IterationFunc[T any] func(node T) error

// IterationDirection 迭代方向
type IterationDirection int8

const (
	StartToFinish IterationDirection = iota // 从头部到尾部
	FinishToStart IterationDirection = iota // 从尾部到头部
)

type IList[T any] interface {
	Size() uint                                                                           // 长度
	IsEmpty() bool                                                                        // IsEmpty 是否为空
	IsFull() bool                                                                         // IsFull 是否满
	HeadPush(val T) error                                                                 // 头插
	HeadPop() (T, error)                                                                  // 头出
	TailPush(val T) error                                                                 // 尾插
	TailPop() (T, error)                                                                  // 尾出
	HeadPeek() (T, error)                                                                 // 头部元素
	TailPeek() (T, error)                                                                 // 尾部元素
	Iteration(iterationFunc IterationFunc[T], directionParam ...IterationDirection) error // 迭代
	ToArray() []T                                                                         // 转换为数组
	stack.IStack[T]                                                                       // 实现栈接口
	queue.IQueue[T]                                                                       // 实现队列接口
}
