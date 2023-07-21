package hash_table

import (
	"fmt"
	"strconv"
)

// This is absolutely not anywhere near production code. It's a really simple
// example which shouldn't be used in real code. Use https://pkg.go.dev/hash instead.

// Hasher represents the strategy the hash table should use to insert items.
type Hasher interface {
	ComputeHash(string) int
}

type Hash struct {
	Key int
}

// AdditiveHasher uses a very naive algorithm that sums the int values of each character.
// Pros:
// - Stable
// - Efficient
// Cons:
// - Not Uniform: "foo" == "oof"
// - Not Secure
type AdditiveHasher struct{}

func (ah AdditiveHasher) ComputeHash(s string) int {
	key := 0

	for _, c := range s {
		key += int(c)
	}

	return key
}

// FoldingHasher creates a hash by folding a moving window of 4 characters together and then
// adding to a rolling value.
// Pro:
// - Stable
// - Efficient
// - Uniform: "foo" != "oof"
// Cons:
// - Not Secure
type FoldingHasher struct{}

func (ah FoldingHasher) ComputeHash(s string) int {
	idx := 0
	key := 0

	for idx < len(s)-4 {
		window := s[idx : idx+4]
		var t string
		for _, c := range window {
			t = fmt.Sprintf("%d%s", int(c), t)
		}
		i, _ := strconv.Atoi(t)
		key += i
		idx += 4
	}

	return key
}

type HashTable struct {
	hasher   Hasher
	table    map[int]string
	capacity int
}

// new creates an AdditiveHash containing the key for the string s provided
func new(hsr Hasher, size int) HashTable {
	hashTable := make(map[int]string, size)
	return HashTable{
		table:    hashTable,
		capacity: size,
		hasher:   hsr,
	}
}

// Add uses the hash to insert the item into the hash table. There may be
// collisions because the table could be relatively small and in that case
// the new item replaces the old.
func (ht *HashTable) Add(item string) {
	hash := Hash{}
	hash.Key = ht.hasher.ComputeHash(item)
	idx := hash.Key % ht.capacity

	ht.table[idx] = item
}

// Get obtains the item from the hash table by computing its hash to get the location.
func (ht *HashTable) Get(item string) string {
	hash := Hash{}
	hash.Key = ht.hasher.ComputeHash(item)
	idx := hash.Key % ht.capacity

	return ht.table[idx]
}
