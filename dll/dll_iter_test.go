package dll

import (
	"testing"
)

func TestIterator(t *testing.T) {
	int_list := NewDeque[int]()
	for i := 1; i < 3; i++ {
		int_list.Append(i)
		int_list.Push(i * 7)
	}

	count := 0
	it1 := int_list.Iterator(0)
	for ; it1.IsValid(); it1.Next() {
		count++
	}

	if it1.Next() {
		t.Errorf("You can't go anymore back we are past the tail of the list")
	}

	if count != 4 {
		t.Errorf("Error There are 4 elements in the iterator; we should have traversed all of them")
	}
	count = 0
	it2 := int_list.Iterator(1)
	for ; it2.IsValid(); it2.Prev() {
		count++
	}

	if it2.Prev() {
		t.Errorf("You can't go anymore back we are past the head of the list")
	}

	if count != 4 {
		t.Errorf("Error There are 4 elements in the iterator; we should have traversed all of them")
	}

	it3 := int_list.Iterator(0)
	if !it3.HasNext() {
		t.Errorf("Incorrect Iterator Construction")
	}

	if it3.HasPrev() {
		t.Errorf("Incorrect Iterator Construction")
	}

	if it3.GetPayload() != 14 {
		t.Errorf("Incorrect Iterator Construction")
	}
}
