package stack

import (
	"fmt"
	"testing"
)

type TNode struct {
	Val int
}

func TestArrayStack_Push(t *testing.T) {
	stackCap := 5
	stack := NewArrayStack[*TNode](uint(stackCap))
	for i := 0; i < stackCap; i++ {
		err := stack.Push(&TNode{Val: i + 1})
		if err != nil {
			t.Error(fmt.Sprintf("stack.push error: %s", err.Error()))
			return
		}
	}
	for !stack.IsEmpty() {
		node, err := stack.Pop()
		if err != nil {
			t.Error(fmt.Sprintf("stack.pop error: %s", err.Error()))
			return
		}
		fmt.Println(node.Val)
	}
}

func TestLinkStack_Push(t *testing.T) {
	stack := NewLinkStack[int]()
	for i := 0; i < 5; i++ {
		err := stack.Push(i)
		if err != nil {
			t.Error(fmt.Sprintf("stack.push error: %s", err.Error()))
			return
		}
	}

	for !stack.IsEmpty() {
		item, err := stack.Pop()
		if err != nil {
			t.Error(fmt.Sprintf("stack.pop error: %s", err.Error()))
			return
		}
		fmt.Println(item)
	}

	fmt.Println("size: ", stack.Size())
	fmt.Println(stack.Pop())
}
