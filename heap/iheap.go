package heap

import "errors"

var IsEmpty = errors.New("heap is empty")

// IHeap 堆接口定义
type IHeap[T comparable] interface {
	Push(node T)      // 插入元素
	Pop() (T, error)  // 移除优先级最高的元素并返回
	Peek() (T, error) // 获取优先级最高的元素
	Clear()           // 清空
	Size() int        // 获取有效元素的个数
	IsEmpty() bool    // 检测优先级队列是否为空
}

// IComparable 比较器
type IComparable[T any] func(node1 T, node2 T) int
