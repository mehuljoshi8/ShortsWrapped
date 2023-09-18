package ht

import "recipeBot/dll"

// A HashTable is an array of buckets,
// where each bucket is a linked list of HTKeyValue structs.
type HashTable struct {
    num_buckets     uint64
    num_elements    uint64
    buckets         []*dll.LinkedList;
}

type HTKey_t = uint64
type HTValue_t = interface{}
type HTKeyValue_t struct {
    key     HTKey_t
    value   HTValue_t
}

// Internal hash function used to map from HTKey_t keys
// to a bucket number.
func (ht *HashTable) HashKeyToBucketNum(key HTKey_t) uint64 {
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
// HTKey_t FVNHash64(buffer, int len)
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
func (ht *HashTable) Resize() {
}

// Returns the number of elements in the hash table.
func (ht *HashTable) GetNumElements() uint64 {
    return ht.num_elements
}

// TODO: Implement Insert.
// Inserts a (key,value) pair into the HashTable.
//
// Arguments:
// - table: the HashTable to insert into.
// - newkeyvalue: the HTKeyValue_t to insert into the table.
// - oldkeyval: if the key in newkeyvalue is already present
//   in the HashTable, that old (key,value) is replaced with
//   newkeyvalue.  In that case, the old (key,value) is returned via
//   this return parameter to the caller.  It's up to the caller
//   to free any allocated memory associated with oldkeyvalue->value.
//
// Returns:
//  - false: if the newkeyvalue was inserted and there was no
//    existing (key,value) with that key.
//  - true: if the newkeyvalue was inserted and an old (key,value)
//    with the same key was replaced and returned through
//    the oldkeyval return parameter.  In this case, the caller assumes
//    ownership of oldkeyvalue.
func (ht *HashTable) Insert(newkv HTKeyValue_t) (bool, *HTKeyValue_t) {
    return false, nil
}


// TODO: Implement Find
// Looks up a key in the HashTable, and if it is present, returns the
// (key,value) associated with it.
//
// Arguments:
// - table: the HashTable to look in.
// - key: the key to look up.
// - keyvalue: if the key is present, a copy of the (key,value) is
//   returned to the caller via this return parameter.  Note that the
//   (key,value) is left in the HashTable, so it is not safe for the
//   caller to free keyvalue->value.
//
// Returns:
//  - false: if the key wasn't found in the HashTable.
//  - true: if the key was found, and therefore the associated (key,value)
//    was returned to the caller via that keyvalue return parameter.
func (ht *HashTable) Find(key HTKey_t) (bool, *HTKeyValue_t) {
    return false, nil
}

// TODO: Implement Remove
// Removes a (key,value) from the HashTable and returns it to the
// caller.
//
// Arguments:
// - table: the HashTable to look in.
// - key: the key to look up.
// - keyvalue: if the key is present, a copy of (key,value) is returned
//   to the caller via this return parameter and the (key,value) is
//   removed from the HashTable.  Note that the caller is responsible
//   for managing the memory associated with keyvalue->value from
//   this point on.
//
// Returns:
//  - false: if the key wasn't found in the HashTable.
//  - true: if the key was found, and therefore (a) the associated
//    (key,value) was returned to the caller via that keyvalue return
//    parameter, and (b) that (key,value) was removed from the
//    HashTable.
func (ht *HashTable) Remove(key HTKey_t) (bool, *HTKeyValue_t) {
    return false, nil
}
