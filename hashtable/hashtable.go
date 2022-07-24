package hashtable

import (
	"bytes"
	"encoding/gob"
	"hash/maphash"
	"myhashtable/list"
)

const (
	defaultCapacity   int     = 10
	defaultLoadFactor float64 = 0.75
)

type Hashtable struct {
	cap           int
	len           int
	reservedLists int
	loadFactor    float64
	table         []*list.LinkedList
	hashFunc      func(val any) uint64
}

type Pair struct {
	key   any
	value any
}

// GetValue returns just value of the key
func GetValue(pair *Pair) any {
	return pair.value
}

// newHashFunc creates and returns hash function
func newHashFunc() func(key any) uint64 {
	h := maphash.Hash{}
	h.SetSeed(maphash.MakeSeed()) // like salt in passwords, not necessary
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	return func(key any) uint64 {
		h.Reset()
		buf.Reset()
		_ = enc.Encode(key)
		_, _ = h.Write(buf.Bytes())
		sum := h.Sum64()
		return sum
	}
}

// newTable creates new array of lists and returns it
func newTable(capacity int) []*list.LinkedList {
	table := make([]*list.LinkedList, capacity)
	for i := 0; i < capacity; i++ {
		table[i] = list.NewLinkedList()
	}
	return table
}

// NewHashTable creates new empty hashtable
func NewHashTable() *Hashtable {
	return &Hashtable{
		cap:           defaultCapacity,
		len:           0,
		reservedLists: 0,
		loadFactor:    defaultLoadFactor,
		table:         newTable(defaultCapacity),
		hashFunc:      newHashFunc(),
	}
}

// NewPair creates and return a pointer to new empty key-value pair
func NewPair(key, value any) *Pair {
	return &Pair{
		key:   key,
		value: value,
	}
}

// Get searches key-value pair in all hashtable and returns it if is found, otherwise returns nil
func (h *Hashtable) Get(key any) *Pair {
	hashSum := h.hashFunc(key) % uint64(h.cap)
	linkedlist := h.table[hashSum]
	for node := linkedlist.Front(); node != nil; node = node.Next() { // rename l
		if pair, ok := node.Val.(*Pair); ok && pair.key == key {
			return pair
		}
	}
	return nil
}

// Set creates new key-value pair if it doesn't exists, otherwise rewrites the value of the key
func (h *Hashtable) Set(key, value any) {
	if pair := h.Get(key); pair != nil {
		pair.value = value
	} else {
		pair = NewPair(key, value)
		index := h.hashFunc(key) % uint64(h.cap)
		_ = h.table[index].PushBack(pair)
		h.reservedLists++
		if float64(h.reservedLists)/float64(h.cap) > h.loadFactor {
			h.Rehash()
		}
	}
}

// Remove removes key-value pair
func (h *Hashtable) Remove(key any) {
	hashSum := h.hashFunc(key) % uint64(h.cap)
	linkedlist := h.table[hashSum]
	_ = linkedlist.Remove(key)
	h.reservedLists--
}

// Rehash creates new double-sized hashtable and moves all data from old one to new one, counting new hash
func (h *Hashtable) Rehash() {
	t := newTable(h.cap * 2)
	h.reservedLists = 0
	for i := 0; i < h.cap; i++ {
		l := h.table[i]
		for node := l.Front(); node != nil; node = node.Next() {
			pair := node.Val.(*Pair)
			index := h.hashFunc(pair.key) % uint64(h.cap*2)
			if t[index].IsEmpty() {
				h.reservedLists++
			}
			_ = t[index].PushBack(pair)
		}
	}
	h.cap *= 2
	h.table = t
	if float64(h.reservedLists)/float64(h.cap) > h.loadFactor {
		h.Rehash()
	}
}
