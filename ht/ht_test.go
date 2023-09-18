package ht

import "testing"

func TestAllocateHashTable(t *testing.T) {
    ht := AllocateHashTable(10)
    if ht.num_buckets != 10 {
        t.Errorf("error expected 10 buckets got %d", ht.num_buckets)
    }
    
    if ht.num_elements != 0 {
        t.Errorf("allocate hashtable should not create any elems")
    }
    
    if len(ht.buckets) != 10 {
        t.Errorf("did not create 10 buckets for buckets field in hashtable")
    }
}


