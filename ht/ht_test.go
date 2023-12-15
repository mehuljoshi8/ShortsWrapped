package ht

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	ht := NewHashTable[int, string]()
	if ht.num_elements != 0 {
		t.Errorf("ht is an empty hash table we should not have elements in it")
	}

	if ht.num_buckets != INIT_NUM_BUCKETS {
		t.Errorf("incorrectly constructed a hash table; initially all hash tables have 10 buckets")
	}
}
