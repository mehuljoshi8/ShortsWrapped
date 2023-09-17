package ht

import "recipeBot/dll"

// A HashTable is an array of buckets,
// where each bucket is a linked list. Of HTKeyValue structs.
type HashTable struct {
    num_buckets     uint64
    num_elements    uint64
    buckets         []dll.LinkedList;
}

// The HashTable Iterator
/*
type HTIterator struct {
    ht              *HashTable
    bucket_idx      uint64
    bucket_it       *dll.LLIter 
}
*/

// int HashKeyToBucketNum(ht *HashTable, HTKey_t key)

type HTKey_t = uint64
type HTValue_t = interface{}
type HTKeyValue_t struct {
    key     HTKey_t
    value   HTValue_t
}

// TODO: Implement FVN hash
// FNV hash implementation.
//
// Customers can use this to hash an arbitrary sequence of bytes into
// a 64-bit key suitable for using as a hash key.  If you're curious, you
// can read about FNV hashing here:
//     http://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function
//
// Arguments:
// - buffer: a pointer to a len-size buffer of unsigned chars.
// - len: how many bytes are in the buffer.
//
// Returns:
// - a nicely distributed 64-bit hash value suitable for
//   use in a HTKeyValue_t.
// HTKey_t FVNHash64(buffer, int len)


// TODO: Implement HashTableAllocate(num_buckets)
// Allocates and returns a new HashTable.
// Argugments:
//  - num_buckets: the number of buckets the hash table should initally
//      contain; MUST be greater than zero.
// Returns nil on error, non-null on success.
func HashTableAllocate(num_buckets uint64) *HashTable {
    return nil
}

// TODO: Implement NumElements
// Figure out the number of elements in the hash table.
//
// Arguments:
//
// - table:  the table to query
//
// Returns:
//
// - table size (>=0); note that this is an unsigned 64-bit integer.
func (ht *HashTable) GetNumElements() int {
    return 0
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
