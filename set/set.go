package set

type ISet[T any] interface {
	Add(item T)
	Remove(item T)
	Contains(item T) bool
}
