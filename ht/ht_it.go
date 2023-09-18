package ht

import "recipeBot/dll"

// The HashTable Iterator
type HTIterator struct {
    ht              *HashTable
    bucket_idx      uint64
    bucket_it       *dll.LLIter
}


// and the functions with it
