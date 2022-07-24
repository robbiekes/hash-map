package main_test

import (
	"fmt"
	"myhashtable/hashtable"
	"myhashtable/list"
	"testing"
)

func TestLinkedList_Remove(t *testing.T) {
	l := list.NewLinkedList()
	_ = l.PushBack(1)
	_ = l.PushBack(2)
	_ = l.PushBack(3)
	_ = l.Remove(15)
	printList(l)
	fmt.Println("len: ", l.Len())
}

func printList(l *list.LinkedList) {
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Print(n.Val, " ")
	}
	fmt.Println()
}

func TestHashtable_Insert(t *testing.T) {
	ht := hashtable.NewHashTable()
	ht.Set("a", "apple")
	ht.Set("a", "ant")
	ht.Set(1, 2)
	c := make(chan struct {
		a int
		b int
	})
	s := struct {
		a, b           int
		etototsamymish string
	}{1, 2, "mav"}
	ht.Set(s, c)
	pair := ht.Get("a")
	fmt.Println(pair)
}
