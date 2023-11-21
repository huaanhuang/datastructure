package heap

import (
	"fmt"
	"testing"
)

type Node struct {
	Val int
}

func Test_Heap(t *testing.T) {
	h := NewHeap[*Node](func(node1 *Node, node2 *Node) int {
		return node2.Val - node1.Val
	})

	h.Push(&Node{Val: 8})
	h.Push(&Node{Val: 1})
	h.Push(&Node{Val: 3})
	h.Push(&Node{Val: 5})
	h.Push(&Node{Val: 2})
	h.Push(&Node{Val: 1})
	h.Push(&Node{Val: 5})
	h.Push(&Node{Val: 6})
	h.Push(&Node{Val: 9})
	h.Push(&Node{Val: 6})

	fmt.Println(h.Size())
	peek, _ := h.Peek()
	fmt.Println(peek.Val == 1)
	fmt.Println("Peek: ", peek)

	for !h.IsEmpty() {
		fmt.Println(h.Pop())
	}

	pop, err := h.Pop()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("pop: ", pop)
}
