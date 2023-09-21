package ht

import "recipeBot/dll"

// The HashTable Iterator
type HTIterator struct {
    ht              *HashTable
    bucket_idx      int64
    bucket_it       *dll.LLIter
}

// Manufactures an iterator for the hashtable.
// If there are elements in the hash table, then the
// iterator points to the "first" one.
// Returns a newly allocated iterator, which may
// be invaid or "past the end" if the table
// can't be iterated through (empty).
func (ht *HashTable) Iterator() *HTIterator {
    htIter := new(HTIterator)
    if ht.num_elements == 0 {
        htIter.ht = ht
        htIter.bucket_idx = -1
        htIter.bucket_it = nil
        return htIter
    }

    htIter.ht = ht
    for i := 0; i < int(ht.num_buckets); i++ {
        if ht.buckets[i].GetSize() > 0 {
            htIter.bucket_idx = int64(i)
            break
        }
    }

    htIter.bucket_it = ht.buckets[htIter.bucket_idx].Iterator()
    return htIter
}

// Tests to see if the iterator is pointing to a valid element
func (htIter *HTIterator) IsValid() bool {
    if htIter.bucket_it == nil {
        return false
    }

    return htIter.bucket_it.IsValid()
}

// Advance the iterator to the next element in the table.
// Returns true if we successfully advanced the iterator.
// False if we can't advance it (i.e. it's past the end)
func (htIter *HTIterator) Next() bool {
    return false
}

// Returns a copy of what the iteartor is currently pointing to 
// or false, nil if the iterator is "past the end"
func (htIter *HTIterator) Get() (bool, *HTKeyValue_t) {
    return false, nil
}

// Returns a copy of (key, value) that the iterator is currently pointing
// at and removes the key, value from the hashtable and advances the iterator
// to the next element.
// Returns false if the remove operation is on an empty table
// and true if the deletion was successful
func (htIter *HTIterator) Remove() (bool, *HTKeyValue_t) {
    return false, nil
}
