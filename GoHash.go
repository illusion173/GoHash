package main

import (
	"fmt"
	"hash/fnv"
)

// This is the initial size of the hashmap
var (
	num_of_pairs    int    = 0
	size_of_table   int    = 10
	size_of_table_u uint32 = 10
	key_slice       []string
	value_slice     []uint32
)

// Begin here
func main() {
	fmt.Println("Beginning GoHash Program")
	// Initial data to hash
	keys_arr := [5]string{"John", "Jerry", "A.J", "Gabe", "Hannah"}
	value_arr := [5]uint32{20, 21, 54, 23, 33}
	// Initially of course both slices will not have anything
	// These will operate as the arrays for hashed data
	key_slice = make([]string, size_of_table)
	value_slice = make([]uint32, size_of_table)
	// Physical inputting of data for processing
	for i := 0; i < len(keys_arr); i++ {
		put(keys_arr[i], value_arr[i])
	}

	text, worked := get("John")
	fmt.Println(text)
	fmt.Println(worked)
}

func get(key any) (uint32, bool) {
	if key == nil {
		fmt.Println("Key cannot be nil in get operation")
		return 0, false
	}

	i := hashKey(key)

	for key_slice[i] != "" {
		if key_slice[i] == key {
			return value_slice[i], true
		}
	}
	return 0, false
}

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
		// Double table size if 50% full.
		size_of_table *= 2
		size_of_table_u *= 2
		key_slice = resizeKey()
		value_slice = resizeValue()
	}

	i := hashKey(key)

	for key_slice[i] != "" {
		if key_slice[i] == key {
			value_slice[i] = value.(uint32)
			return
		}
		i = ((i + 1) % size_of_table_u)
	}

	key_slice[i] = key.(string)
	value_slice[i] = value.(uint32)
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

func hashKey(key any) uint32 {
	var h uint32 = 0

	switch key.(type) {
	case string:
		h = hashCode(key.(string))
	default:
		h = key.(uint32)
	}
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)
	return (h & (size_of_table_u - 1))
}
