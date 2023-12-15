/* Author: Mehul Joshi
 * Tests for a Deque to ensure
 * that it's correctly built :)
 */
package dll

import (
	"math/rand"
	"testing"
)

func TestEmptyConstructor(t *testing.T) {
	var int_list = NewDeque[int]()
	if int_list.size != 0 {
		t.Errorf("list has a non-zero initial size after construction: %q", int_list.size)
	}

	if int_list.head != nil || int_list.tail != nil {
		t.Errorf("the list's head and tail should be nil initially")
	}
}
func TestPush(t *testing.T) {
	var int_list = NewDeque[int]()
	for i := 0; i < 1000; i++ {
		int_list.Push(i)
		if int_list.size != (i+1) || int_list.head.payload != i {
			t.Errorf("in correct updating of values new values should be added to the head of the list")
		}
	}
}

func TestPushPop(t *testing.T) {
	var int_list = NewDeque[int]()
	popped, value := int_list.Pop()
	if popped || value != zero[int]() {
		t.Errorf("Popping on an empty list should not be allowed: %q", value)
	}

	for i := 0; i < 100; i++ {
		int_list.Push(i)
	}

	for i := 99; i >= 0; i-- {
		popped, value := int_list.Pop()
		if !popped || value != i {
			t.Errorf("popping values of the linked list should update correctly")
			t.Errorf("\tExpected: %q", i)
			t.Errorf("\tActual: %q", value)
		}
	}
}

func TestPushPopPeek(t *testing.T) {
	int_list := NewDeque[int]()
	var peeked, value = int_list.Peek()
	if peeked || value != zero[int]() {
		t.Errorf("Peek on an empty list should always return false and a zero value")
	}

	for i := 0; i < 100; i++ {
		int_list.Push(i)
		peeked, value := int_list.Peek()
		if !peeked || value != i {
			t.Errorf("Peeking on a non-empty list should always return true")
		}

		if int_list.Size() != (i + 1) {
			t.Errorf("Size is not being updated properly")
		}
	}

	for i := 99; i >= 0; i-- {
		peeked, value := int_list.Peek()
		int_list.Pop()
		if !peeked || value != i {
			t.Errorf("peeking values of the linked list should update correctly")
			t.Errorf("\tExpected: %q", i)
			t.Errorf("\tActual: %q", value)
		}
	}
}

func TestAppendSlicePeekBack(t *testing.T) {
	int_list := NewDeque[int]()
	var peeked, value = int_list.PeekBack()

	if peeked || value != zero[int]() {
		t.Errorf("Peek on an empty list should always return false and a zero value")
	}

	peeked, value = int_list.Slice()
	if peeked || value != zero[int]() {
		t.Errorf("Slice on an empty list should always return false and a zero value")
	}

	for i := 0; i < 100; i++ {
		int_list.Append(i)
		peeked, value = int_list.PeekBack()
		if !peeked || value != i {
			t.Errorf("Peeking on a non-empty list should always return true")
		}

		if int_list.Size() != (i + 1) {
			t.Errorf("Size is not being updated properly")
		}
	}

	for i := 99; i >= 0; i-- {
		peeked, value = int_list.PeekBack()
		int_list.Slice()
		if !peeked || value != i {
			t.Errorf("peeking values of the linked list should update correctly")
			t.Errorf("\tExpected: %q", i)
			t.Errorf("\tActual: %q", value)
		}
	}
}

// Comparator function for dll
func comp_fn(p1, p2 interface{}) int {
	return p1.(int) - p2.(int)
}

func TestSort(t *testing.T) {
	// create a linked list with random elems
	ll := NewDeque[int]()

	ll.Sort(comp_fn)

	size := rand.Intn(10000)
	for i := 0; i < size; i++ {
		ll.Push(rand.Intn(100))
	}

	ll.Sort(comp_fn)
	b, v := ll.Pop()
	count := 0
	if b {
		count = 1
	}
	for ll.Size() > 0 {
		_, v2 := ll.Pop()
		if v2 < v {
			t.Errorf("error in sorting %v >= %v", v2, v)
		}
		v = v2
		count++
	}

	if count != size {
		t.Errorf("elements got lost after sorting: count = %v, size = %v", count, size)
	}
}
