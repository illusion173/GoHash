package main

import (
	"fmt"
	"hash/fnv"
)

var num_of_pairs int = 0

// This is the initial size of the hashmap
var (
	size_of_table   int    = 10
	size_of_table_u uint32 = 10
	key_slice       []string
	value_slice     []uint32
)

// Begin here
func main() {
	fmt.Println("Beginning GoHash Program")
	// unfortunate hard coding gotta learn generics for slices later
	// Initially of course both slices will not have anything
	key_slice = make([]string, size_of_table)
	value_slice = make([]uint32, size_of_table)
	var test uint32 = 0
	for i := 0; i <= 20; i++ {
		put("test", test)
	}
}

// void
func put(key any, value any) {
	if key == nil {
		fmt.Println("Invalid Input")
		fmt.Println("Key cannot be null")
		// cancels insert
		return
	}

	if value == nil {
		fmt.Println("Invalid Input")
		fmt.Println("Warning value is nil with key", key)
		// cancels insert
		return
	}

	if num_of_pairs == (size_of_table / 2) {
		// increase the size of the table prior
		size_of_table *= 2
		size_of_table_u *= 2
		key_slice = resizeKey()
		value_slice = resizeValue()
	}

	key_slice[0] = key.(string)
	value_slice[0] = value.(uint32)
	num_of_pairs += 1
}

func resizeKey() []string {
	newArr := make([]string, size_of_table)
	copy(newArr, key_slice)
	return newArr
}

func resizeValue() []uint32 {
	newArr := make([]uint32, size_of_table)
	copy(newArr, value_slice)
	return newArr
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
	return h & (size_of_table_u - 1)
}
