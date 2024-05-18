package set

func NewSet[T comparable]() ISet[T] {
	return &setMap[T]{
		data: make(map[T]struct{}),
	}
}

type setMap[T comparable] struct {
	data map[T]struct{}
}

// Add 添加元素
func (obj *setMap[T]) Add(item T) {
	obj.data[item] = struct{}{}
}

// Remove 删除元素
func (obj *setMap[T]) Remove(item T) {
	delete(obj.data, item)
}

func (obj *setMap[T]) Contains(item T) bool {
	_, ok := obj.data[item]
	return ok
}
