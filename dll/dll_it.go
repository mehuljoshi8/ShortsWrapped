package dll

// A Linked List Iterator.
type LLIter struct {
    ll *LinkedList
    node *DLLNode
}

// Manufacture an iteartor for the list.
// Usage: ll.Iterator()
// Returns:
//  - a newly-allocated iterator, which may be invalid
//      or "past the end" if the list can't be iterated through (empty).
func (ll *LinkedList) Iterator() *LLIter {
    ll_iter := new(LLIter)
    ll_iter.ll = ll
    ll_iter.node = ll.head
    return ll_iter
}

// Tests to see whether the iterator is pointing at a valid element.
// Returns:
//  - true: if iter is not past the end of the list.
//  - false: if iter is past the end of the list.
func (ll_iter *LLIter) IsValid() bool {
    return ll_iter.node != nil
}

// Advance the iterator, i.e. move to the next node of int the list.
// The passed-in iterator must be valid (eg, not "past the end").
// Returns:
//  - true: if the iterator has been advanced to the next node.
//  - false: if the iterator is no longer valid after the
//           advancing has completed (eg, it's now "past the end").
func (ll_iter *LLIter) Next() bool {
    if ll_iter.IsValid() {
        ll_iter.node = ll_iter.node.next
    }
    return ll_iter.IsValid()
}

// Returns the payload of the list node that the iterator currently points
// at. The passed-in iterator must be valid (eg, not "passed the end").
func (ll_iter *LLIter) Get() LLPayload_t {
    return ll_iter.node.payload
}

// Remove the node the iterator is pointing to.  After deletion, the iterator
// may be in one of the following three states:
// - if there was only one element in the list, the iterator is now invalid
//   and cannot be used.  In this case, the caller is recommended to free
//   the now-invalid iterator.
// - if the deleted node had a successor (ie, it was pointing at the tail),
//   the iterator is now pointing at the successor.
// - if the deleted node was the tail, the iterator is now pointing at the
//    predecessor.
//
// The passed-in iterator must be valid (eg, not "past the end").
//
// Arguments:
// - iter:  the iterator to delete from.
// - payload_free_function: invoked to free the payload.
//
// Returns:
// - false if the deletion succeeded, but the list is now empty.
// - true if the deletion succeeded, and the list is still non-empty.
func (ll_iter *LLIter) Remove() (bool, LLPayload_t) {
    if ll_iter.node == nil || ll_iter.ll == nil {
        // this is an invalid iterator
        return false, nil
    }

    var p LLPayload_t
    if ll_iter.ll.GetSize() == 1 {
        _, p = ll_iter.ll.Pop()
        ll_iter.node = nil
        return true, p
    }

    if ll_iter.node == ll_iter.ll.tail {
        // have to assign ll_iter.node to the previous node
        ll_iter.node = ll_iter.node.prev
        _, p = ll_iter.ll.Slice()
    } else if ll_iter.node == ll_iter.ll.head {
        ll_iter.node = ll_iter.node.next
        _, p = ll_iter.ll.Pop()
    } else {
        // have to do middle of the list random deletion
        p = ll_iter.node.payload
        ll_iter.node = ll_iter.node.next
        ll_iter.node.prev.next = nil
        ll_iter.node.prev.prev.next = nil
        ll_iter.node.prev = ll_iter.node.prev.prev
        ll_iter.node.prev.next = ll_iter.node
        ll_iter.ll.size--
    }

    return true, p
}

// Rewind an iterator to the front of its list.
func (ll_iter *LLIter) Rewind() {
    ll_iter.node = ll_iter.ll.head

}
