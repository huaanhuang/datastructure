package list

import (
	"fmt"
	"github.com/huaanhuang/datastructure/queue"
	"github.com/huaanhuang/datastructure/stack"
	"testing"
)

func Test_List(t *testing.T) {
	// 栈
	var s stack.IStack[int] = NewLinkList[int]()
	// 队列
	var q queue.IQueue[int] = NewLinkList[int]()
	// 双端链表
	var l IList[int] = NewLinkList[int]()
	fmt.Println(s, q, l)
}

func Test_Stack(t *testing.T) {
	testArr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var s stack.IStack[int] = NewLinkList[int]()
	for _, item := range testArr {
		err := s.Push(item)
		if err != nil {
			t.Error(err)
			return
		}
	}
	for !s.IsEmpty() {
		pop, err := s.Pop()
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(pop)
	}
	fmt.Println("size: ", s.Size())
	fmt.Println("isEmpty: ", s.IsEmpty())
	peek, err := s.Peek()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("peek: ", peek)
}

func Test_Queue(t *testing.T) {
	testArr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var s queue.IQueue[int] = NewLinkList[int]()
	for _, item := range testArr {
		err := s.Put(item)
		if err != nil {
			t.Error(err)
			return
		}
	}
	for !s.IsEmpty() {
		get, err := s.Get()
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(get)
	}
	fmt.Println("size: ", s.Qsize())
	fmt.Println("isEmpty: ", s.IsEmpty())
}
