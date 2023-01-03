package main

import (
	"fmt"
	"hash/fnv"
)

var num_of_pairs int = 0

// This is the initial size of the hashmap
var (
	size_of_table int = 10
	key_slice     []string
	value_slice   []uint64
)

// Begin here
func main() {
	fmt.Println("Beginning GoHash Program")
	// unfortunate hard coding gotta learn generics for slices later
	// Initially of course both slices will not have anything
	key_slice = make([]string, num_of_pairs, size_of_table)
	value_slice = make([]uint64, num_of_pairs, size_of_table)
}

// void
func insert(key any, value any) {
	if key == nil {
		panic("Key cannot be null")
	}

	if value == nil {
		fmt.Println("Warning value is nil with key", key)
	}
}

func hashCode(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func hash(key any) uint32 {
	var h uint32 = 0

	switch key.(type) {
	case string:
		h = hashCode(key.(string))
	default:
		h = key.(uint32)
	}
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)
	return h & (uint32(size_of_table) - 1)
}
