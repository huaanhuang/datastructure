package stack

// IStack 栈接口
type IStack[T any] interface {
	// Push 入栈
	Push(item T) error
	// Pop 出栈
	Pop() (item T, err error)
	// IsEmpty 栈是否为空
	IsEmpty() bool
	// Size 栈当前大小
	Size() uint
	// Peek 返回栈顶元素
	Peek() (item T, err error)
}
