package ht

import "recipeBot/dll"

// A HashTable is an array of buckets,
// where each bucket is a linked list of HTKeyValue structs.
type HashTable struct {
    num_buckets     uint64
    num_elements    uint64
    buckets         []*dll.LinkedList;
}

// type aliases for the HashTable
type HTKey_t = uint64
type HTValue_t = interface{}
// The KeyValue struct we are going to use.
type HTKeyValue_t struct {
    key     HTKey_t
    value   HTValue_t
}

// Internal hash function used to map from HTKey_t keys
// to a bucket number.
func (ht *HashTable) hashKeyToBucketNum(key HTKey_t) uint64 {
    return key % ht.num_buckets
}

// FNV hash implementation.
//
// Customers can use this to hash an arbitrary sequence of bytes into
// a 64-bit key suitable for using as a hash key.  If you're curious, you
// can read about FNV hashing here:
//     http://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function
//
// Arguments:
// - buffer: of bytes that needs to hashed to a key
//
// Returns:
// - a nicely distributed 64-bit hash value suitable for
//   use in a HTKeyValue_t.
func FNVHash64(buffer []byte) HTKey_t {
    const FNV1_64_INIT uint64 = 0xcbf29ce484222325
    const FNV_64_PRIME uint64 = 0x100000001b3
    hval := FNV1_64_INIT
    // hash each octet of the buffer
    for _, b := range buffer {
        // XOR the bottom wtih the current octet.
        hval ^= HTKey_t(b)
        // Multiply by the 64 bit FNV magic prime mod 2 ^64.
        hval *= FNV_64_PRIME
    }
    return hval
}

// Allocates and returns a new HashTable with num_buckets
// number of buckets
func AllocateHashTable(num_buckets uint64) *HashTable {
    var ht *HashTable = new(HashTable)
    ht.num_buckets = num_buckets
    ht.num_elements = 0
    ht.buckets = make([]*dll.LinkedList, num_buckets)
    for i := 0; i < int(num_buckets); i++ {
        ht.buckets[i] = dll.AllocateLinkedList()
    }
    return ht
}

// Grows the hashtable (increases the number of buckets) if
// its load factor has become to high.
// TODO: Complete the implementation of resize
func (ht *HashTable) resize() {
    if ht.num_elements < 3 * ht.num_buckets {
        // no need to resize
        return
    }

    // newht := AllocateHashTable(ht.num_buckets * 3)

    // implement ht iterator to complete the rest of the func.
}

// Returns the number of elements in the hash table.
func (ht *HashTable) GetNumElements() uint64 {
    return ht.num_elements
}

// TODO: Implement Insert.
// Inserts a (key,value) pair into the HashTable (denoted newkv)
// and if the key is already present in the hashtable with a
// different associated value then we return that old HTKeyValue_t
// to the user. if it did not exist then we return nil for the return value.
func (ht *HashTable) Insert(newkv HTKeyValue_t) *HTKeyValue_t {
    // resize if we need to resize
    //ht.resize()
    //bucketIdx := ht.hashKeyToBucketNum(newkv.key)
    // chain is the linked list that we are going to insert
    // newkv into
    //chain := ht.buckets[bucketIdx] 
    //keyValue := new(HTKeyValue_t)
    //keyValue.key = newkv.key
    //keyValue.value = value
    // implement find optional remove...
    return nil
}


// TODO: Implement Find
// Looks up a key in the HashTable, and if it is present, returns the
// (key,value) associated with it otherwise returns nil
func (ht *HashTable) Find(key HTKey_t) *HTKeyValue_t {
    return nil
}

// TODO: Implement Remove
// Removes a (key,value) from the HashTable and returns it to the
// caller. If the key is not present in the HashTable we return nil.
func (ht *HashTable) Remove(key HTKey_t) *HTKeyValue_t {
    return nil
}

