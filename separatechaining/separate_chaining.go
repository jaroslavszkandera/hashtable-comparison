package separatechaining

import (
	"fmt"
)

const (
	INIT_CAPACITY   = 1024
	MIN_LOAD_FACTOR = 0.25
	MAX_LOAD_FACTOR = 3.00
)

type HashTable struct {
	size     uint
	capacity uint
	data     [][]string
	hashFn   func(string) uint
}

func NewHashTable(hashFn func(string) uint) *HashTable {
	return &HashTable{
		size:     0,
		capacity: INIT_CAPACITY,
		data:     make([][]string, INIT_CAPACITY),
		hashFn:   hashFn,
	}
}

func (h *HashTable) CalcLoadFactor() float32 {
	return float32(h.size) / float32(h.capacity)
}

func simpleHash(val string) uint {
	var hash uint
	for i, char := range val {
		hash += uint(char) * uint(i*i*i*i)
	}
	return hash
}

func (h *HashTable) resize(newCapacity uint) {
	newData := make([][]string, newCapacity)
	for _, bucket := range h.data {
		for _, v := range bucket {
			newIndex := h.hashFn(v) % newCapacity
			newData[newIndex] = append(newData[newIndex], v)
		}
	}
	h.data = newData
	h.capacity = newCapacity
}

func (h *HashTable) Insert(val string) bool {
	if h.CalcLoadFactor() > MAX_LOAD_FACTOR {
		h.resize(h.capacity * 2)
	}

	if h.Search(val) {
		return false
	}

	i := h.hashFn(val) % h.capacity
	h.data[i] = append(h.data[i], val)
	h.size++
	return true
}

func (h *HashTable) Delete(val string) bool {
	if h.CalcLoadFactor() < MIN_LOAD_FACTOR &&
		h.capacity > INIT_CAPACITY {
		h.resize(max(h.capacity/2, INIT_CAPACITY))
	}
	index := h.hashFn(val) % h.capacity
	bucket := h.data[index]

	for i, v := range bucket {
		if v == val {
			bucket = append(bucket[:i], bucket[i+1:]...)
			h.data[index] = bucket
			h.size--
			return true
		}
	}
	return false
}

func (h *HashTable) Search(val string) bool {
	i := h.hashFn(val) % h.capacity
	bucket := h.data[i]

	for _, v := range bucket {
		if v == val {
			return true
		}
	}
	return false
}

func (h *HashTable) Print() {
	fmt.Println("Hash map:")
	for i, b := range h.data {
		if len(b) > 0 {
			fmt.Printf("%d: %v\n", i, b)
		}
	}
}
