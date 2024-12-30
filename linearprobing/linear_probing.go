package linearprobing

import (
	"fmt"
)

const (
	INIT_SIZE          = 1024
	MIN_LOAD_FACTOR_LP = 0.25
	MAX_LOAD_FACTOR_LP = 0.75
	TOMBSTONE          = "<<<TOMBSTONE>>>"
)

type HashTable struct {
	size     uint
	capacity uint
	hashFn   func(string) uint
	data     []string
}

func NewHashTable(hashFn func(string) uint) *HashTable {
	return &HashTable{
		size:     0,
		capacity: INIT_SIZE,
		data: make([]string,
			INIT_SIZE),
		hashFn: hashFn,
	}
}

func (h *HashTable) CalcLoadFactor() float32 {
	return float32(h.size) / float32(h.capacity)
}

func (h *HashTable) resize(newCapacity uint) {
	newData := make([]string, newCapacity)
	for _, v := range h.data {
		if v != "" && v != TOMBSTONE {
			newI := h.hashFn(v) % newCapacity
			var i uint = 0
			for newData[newI] != "" {
				i++
				newI = linProbe(i, newCapacity, h.hashFn(v))
			}
			newData[newI] = v
		}
	}
	h.data = newData
	h.capacity = newCapacity
}

// linear probing (Open addressing)
func linProbe(i, capacity, hashVal uint) uint {
	return (hashVal + i) % capacity
}

func (h *HashTable) Insert(val string) bool {
	if h.CalcLoadFactor() > MAX_LOAD_FACTOR_LP {
		h.resize(uint(h.capacity * 2))
	}

	var i uint = 0
	index := h.hashFn(val) % h.capacity
	var index_tombstone uint = 0
	foundTombstone := false

	for h.data[index] != "" {
		if h.data[index] == val {
			return false
		}
		if h.data[index] == TOMBSTONE && !foundTombstone {
			index_tombstone = index
			foundTombstone = true
		}
		i++
		index = linProbe(i, h.capacity, h.hashFn(val))
	}

	if foundTombstone {
		h.data[index_tombstone] = val
	} else {
		h.data[index] = val
	}
	h.size++
	return true
}

func (h *HashTable) Delete(val string) bool {
	if !h.Search(val) {
		return false
	}

	index := h.hashFn(val) % h.capacity
	for i := uint(0); i < h.capacity; i++ {
		if h.data[index] == val {
			break
		}
		index = linProbe(i, h.capacity, h.hashFn(val))
	}
	h.data[index] = TOMBSTONE
	h.size--

	if h.CalcLoadFactor() < MIN_LOAD_FACTOR_LP && h.capacity > INIT_SIZE {
		h.resize(max(h.capacity/2, INIT_SIZE))
	}
	return true
}

func (h *HashTable) Search(val string) bool {
	index := h.hashFn(val) % h.capacity
	for i := uint(0); i < h.capacity; i++ {
		if h.data[index] == "" {
			return false
		}
		if h.data[index] == val {
			return true
		}
		index = linProbe(i, h.capacity, h.hashFn(val))
	}
	return false
}

func (h *HashTable) Print() {
	fmt.Println("Hash map:")
	for i, v := range h.data {
		if v != "" && v != TOMBSTONE {
			fmt.Printf("%d: %v\n", i, v)
		}
	}
}
