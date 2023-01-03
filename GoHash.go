package main

import(
  "hash/fnv"
  "fmt"
) 
var num_of_pairs int = 0
// This is the initial size of the hashmap
var size_of_table int = 10
var key_slice []string
var value_slice []uint64

//Begin here
func main(){
  fmt.Println("Beginning GoHash Program")
  // unfortunate hard coding gotta learn generics for slices later

  // Initially of course both slices will not have anything
  value_slice = make([]uint64, num_of_pairs, size_of_table)
  key_slice = make([]string, num_of_pairs, size_of_table)
  insert(nil, nil)
}
func insert(key any, value any) bool {
  if(key == nil){
    panic("Key cannot be null")
  }

  if(value == nil){
    fmt.Println("Warning value is nil with key", key)
  }

  return true
}
func hashCode(s string) uint64 {
        h := fnv.New64a()
        h.Write([]byte(s))
        return h.Sum64()
}
func hash(key any)(uint64) {
	if value, ok := key.(string); ok {
    var h uint64 = hashCode(value)
    fmt.Println(h)
	}
  var h uint64 = key.(uint64)
  h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) & (h >> 4)
  return h
}
