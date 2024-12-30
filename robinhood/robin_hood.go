package robinhood

import (
	"fmt"
)

const (
	INIT_CAPACITY   = 1024
	MIN_LOAD_FACTOR = 0.25
	MAX_LOAD_FACTOR = 0.75
)

type Elem struct {
	psl  uint // probe sequence length
	data string
}

type HashTable struct {
	size     uint
	capacity uint
	hashFunc func(string) uint
	elems    []Elem
}

func NewHashTable(hashFunc func(string) uint) *HashTable {
	return &HashTable{
		size:     0,
		capacity: INIT_CAPACITY,
		hashFunc: hashFunc,
		elems:    make([]Elem, INIT_CAPACITY),
	}
}

func (h *HashTable) CalcLoadFactor() float32 {
	return float32(h.size) / float32(h.capacity)
}

func (h *HashTable) resize(newCapacity uint) {
	newElems := make([]Elem, newCapacity)
	oldElems := h.elems
	oldCapacity := h.capacity

	h.elems = newElems
	h.capacity = newCapacity
	h.size = 0

	for i := uint(0); i < oldCapacity; i++ {
		if oldElems[i].data != "" {
			h.Insert(oldElems[i].data)
		}
	}
}

func (h *HashTable) Insert(val string) bool {
	if h.CalcLoadFactor() > MAX_LOAD_FACTOR {
		h.resize(h.capacity * 2)
	}

	hash := h.hashFunc(val) % h.capacity

	// start probing for an empty slot
	swapElem := Elem{psl: 0, data: val}
	var probe uint = 0
	for {
		index := (hash + probe) % h.capacity
		if h.elems[index].data == "" {
			h.elems[index] = swapElem
			h.size++
			return true
		}
		if h.elems[index].data == val {
			return false // duplicate
		}
		if h.elems[index].psl < swapElem.psl { // swap
			tmp := h.elems[index]
			h.elems[index] = swapElem
			swapElem = tmp
		}
		swapElem.psl++
		probe++
	}
}

func (h *HashTable) Delete(val string) bool {
	if h.CalcLoadFactor() < MIN_LOAD_FACTOR &&
		h.capacity > INIT_CAPACITY {
		h.resize(max(h.capacity/2, INIT_CAPACITY))
	}

	hash := h.hashFunc(val) % h.capacity
	var foundIndex uint = 0
	for probe := uint(0); probe < h.capacity; probe++ {
		index := (hash + probe) % h.capacity
		if h.elems[index].data == "" ||
			h.elems[index].psl < probe {
			return false
		}
		if h.elems[index].data == val {
			foundIndex = index
			break
		}
	}

	h.elems[foundIndex] = Elem{psl: 0, data: ""}
	probe := uint(1)
	for {
		index := (foundIndex + probe) % h.capacity
		if h.elems[index].data == "" || h.elems[index].psl == 0 {
			break
		}
		prevIndex := (index - 1) % h.capacity
		h.elems[prevIndex] = h.elems[index]
		h.elems[prevIndex].psl--
		h.elems[index] = Elem{psl: 0, data: ""}
		probe++
	}
	h.size--
	return true
}

func (h *HashTable) Search(val string) bool {
	hash := h.hashFunc(val) % h.capacity

	for probe := uint(0); probe < h.capacity; probe++ {
		index := (hash + probe) % h.capacity
		if h.elems[index].data == "" ||
			h.elems[index].psl < probe {
			return false
		}
		if h.elems[index].data == val {
			return true
		}
	}
	return false
}

func (h *HashTable) Print() {
	fmt.Println("Hash map print:\n[i: psl elems]")
	for i, elem := range h.elems {
		if elem.data != "" {
			fmt.Printf("%d: %v %v\n", i, elem.psl, elem.data)
		}
	}
}
