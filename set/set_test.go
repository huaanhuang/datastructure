package set

import (
	"fmt"
	"testing"
)

func Test_SetMap(t *testing.T) {
	// 创建一个新的集合
	s := NewSet[string]()

	// 添加元素
	s.Add("apple")
	s.Add("banana")
	s.Add("cherry")

	// 检查元素是否存在
	fmt.Println(s.Contains("apple"))  // 输出: true
	fmt.Println(s.Contains("orange")) // 输出: false

	// 移除元素
	s.Remove("banana")
	fmt.Println(s.Contains("banana")) // 输出: false
}
