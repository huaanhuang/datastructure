package heap

const (
	defaultHeapCap = 11 // 默认容量
)

// NewHeap 实例话堆
func NewHeap[T comparable](compare IComparable[T], cap ...int) IHeap[T] {
	defaultCap := defaultHeapCap
	if len(cap) > 0 && cap[0] > 0 {
		defaultCap = cap[0]
	}
	return &heap[T]{
		data:    make([]T, defaultCap, defaultCap),
		cap:     defaultCap,
		compare: compare,
	}
}

// heap 堆实现
type heap[T comparable] struct {
	data    []T            // 数据
	cap     int            // 容量
	size    int            // 当前堆大小
	compare IComparable[T] // 比较器
}

func (obj *heap[T]) Push(node T) {
	if obj.size >= obj.cap { // 堆满，扩容到原来大小的2倍
		obj.cap = int(float64(obj.cap) * 1.5)
		newData := make([]T, obj.cap, obj.cap)
		copy(newData, obj.data)
		obj.data = newData
	}
	obj.data[obj.size] = node
	obj.heapInsert(obj.size)
	obj.size++
}

func (obj *heap[T]) Pop() (T, error) {
	if obj.size == 0 {
		return *new(T), IsEmpty
	}
	ans := obj.data[0]
	obj.data[0] = obj.data[obj.size-1]
	obj.size--
	obj.heapify(0)
	return ans, nil
}

func (obj *heap[T]) Peek() (T, error) {
	if obj.size == 0 {
		return *new(T), IsEmpty
	}
	return obj.data[0], nil
}

func (obj *heap[T]) Clear() {
	obj.data = make([]T, defaultHeapCap, defaultHeapCap)
	obj.cap = defaultHeapCap
	obj.size = 0
}

func (obj *heap[T]) Size() int {
	return obj.size
}

func (obj *heap[T]) IsEmpty() bool {
	return obj.size == 0
}

// 从index位置开始insert(向上看)
func (obj *heap[T]) heapInsert(index int) {
	head := (index - 1) / 2
	for obj.compare(obj.data[head], obj.data[index]) > 0 {
		obj.data[head], obj.data[index] = obj.data[index], obj.data[head]
		index = head
		head = (index - 1) / 2
	}
}

// 从index位置开始heapify(向下看)
func (obj *heap[T]) heapify(index int) {
	left := 2*index + 1
	for left < obj.size {
		target := left
		right := 2*index + 2
		// 左右孩子谁大是谁
		if right < obj.size && obj.compare(obj.data[right], obj.data[left]) < 0 {
			target = right
		}
		// 两个孩子都没有干过父节点
		if obj.compare(obj.data[target], obj.data[index]) > 0 {
			break
		}
		// 交换
		obj.data[index], obj.data[target] = obj.data[target], obj.data[index]
		// index来到大的孩子继续往下看
		index = target
		left = 2*index + 1
	}
}
