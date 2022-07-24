package list

type List interface {
	PushFront(val any) error
	PushBack(val any) error
	PopFront(val any) error
	PopBack(val any) error
	Remove(val any) error
}
