package list

import "errors"

type Node struct {
	Val  interface{}
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

var (
	ErrEmptyList     = errors.New("empty list")
	ErrValueNotFound = errors.New("value not found")
)

// NewNode creates new empty node
func NewNode(val any, next *Node) *Node {
	return &Node{
		Val:  val,
		next: next,
	}
}

// NewLinkedList creates new enpty list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		len:  0,
	}
}

// PushFront adds a new node to the front of the list
func (l *LinkedList) PushFront(val any) error {
	if l.head == nil {
		l.head = NewNode(val, nil)
		l.tail = l.head
	} else {
		l.head = NewNode(val, l.head)
	}
	l.len++
	return nil
}

// PushBack adds a new node to the back of the list
func (l *LinkedList) PushBack(val any) error {
	if l.tail == nil {
		l.head = NewNode(val, nil)
		l.tail = l.head
	} else {
		l.tail.next = NewNode(val, nil)
		l.tail = l.tail.next
	}
	l.len++
	return nil
}

// PopFront removes the first node of the list
func (l *LinkedList) PopFront() (any, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}
	value := l.head.Val
	l.head = l.head.next
	l.len--
	return value, nil
}

// PopBack removes the last node of the list
func (l *LinkedList) PopBack() (any, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}
	value := l.tail.Val
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		cur := l.head
		for cur.next != l.tail {
			cur = cur.next
		}
		l.tail = cur
		l.tail.next = nil
	}
	l.len--
	return value, nil
}

// Remove removes a node from the list
func (l *LinkedList) Remove(val any) error {
	if l.head == nil {
		return ErrEmptyList
	}
	if l.head.Val == val {
		l.len--
		if l.head == l.tail {
			l.head = nil
			l.tail = nil
			return nil
		}
		l.head = l.head.next
		return nil
	}
	cur := l.head
	for cur.next != nil {
		if cur.next.Val == val {
			cur.next = cur.next.next
			l.len--
			return nil
		}
		cur = cur.next
	}
	return ErrValueNotFound
}

// Front simply returns the first node of the list
func (l *LinkedList) Front() *Node {
	return l.head
}

// Back simply returns the last node of the list
func (l *LinkedList) Back() *Node {
	return l.tail
}

// Len simply returns length of list (number of nodes)
func (l *LinkedList) Len() int {
	return l.len
}

// IsEmpty just checks if hashtable list is empty
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}
