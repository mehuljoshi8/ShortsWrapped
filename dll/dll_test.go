/* Author: Mehul Joshi
 * Tests for a doubly linked list to ensure
 * that it's correctly built :)
 */
package dll

import (
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
		if int_list.size != (i+1) || int_list.head.val != i {
			t.Errorf("in correct updating of values new values should be added to the head of the list")
		}
	}
}
