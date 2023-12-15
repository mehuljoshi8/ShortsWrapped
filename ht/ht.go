package ht

import "recipeBot/dll"

// The smallest unit of storage in our hash table
// 2^64 total keys in our hash table.
// cause of uint64
type HTKeyValue[K interface{}, V interface{}] struct {
	key   K
	value V
}

// We use this function internally to compute the hash function for a sequence of bytes
func fVNHash64(buffer []byte) uint64 {
	const FNV1_64_INIT uint64 = 0xcbf29ce484222325
	const FNV_64_PRIME uint64 = 0x100000001b3
	hval := FNV1_64_INIT
	for _, b := range buffer {
		hval ^= uint64(b)
		hval *= FNV_64_PRIME
	}
	return hval
}

// Returns the bucket that a particular hash table will be in
func hashKeyToBucketNum(key uint64, num_buckets int) uint64 {
	return key % uint64(num_buckets)
}

// All HashTables start off with 10 buckets
const INIT_NUM_BUCKETS int = 10

// The HashTable is just made up of many different buckets which are each of them linked, lists
type HashTable[K interface{}, V interface{}] struct {
	num_elements int
	num_buckets  int
	buckets      [INIT_NUM_BUCKETS]dll.Deque[HTKeyValue[K, V]]
}

// Constructor: returns a new hash table with 10 buckets and no other elements initially
func NewHashTable[K interface{}, V interface{}]() HashTable[K, V] {
	ht := HashTable[K, V]{}
	ht.num_buckets = INIT_NUM_BUCKETS

	for i := 0; i < INIT_NUM_BUCKETS; i++ {
		ht.buckets[i] = dll.NewDeque[HTKeyValue[K, V]]()
	}

	return ht
}

func (ht *HashTable[K, V]) Insert(keyValue HTKeyValue[K, V])

// Returns the number of elements in the HashTable.
func (ht *HashTable[K, V]) Size() int {
	return ht.num_elements
}
