package ht

type HashTable struct {
    num_buckets     uint64
    num_elements    uint64
    buckets         []LinkedList*
}

type HTIterator struct {
    ht              *Hashtable
    bucket_idx      uint64
    bucket_it       *LLIterator 
}

// int HashKeyToBucketNum(ht *HashTable, HTKey_t key)

