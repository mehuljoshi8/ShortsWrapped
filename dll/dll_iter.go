package dll

// iterator definition for a DLL
type DQIter[T interface{}] struct {
	node *dLLNode[T]
}

// Constructs a new DQIter from a deque. pos = 0 -> head
// pos = 1 -> tail
func (dll *Deque[T]) Iterator(pos int) DQIter[T] {
	var startNode *dLLNode[T]
	if pos == 0 {
		startNode = dll.head
	} else {
		startNode = dll.tail
	}

	return DQIter[T]{
		startNode,
	}
}

// Checks to see if we have a next reference
func (iter *DQIter[T]) HasNext() bool {
	return iter.node != nil && iter.node.next != nil
}

// Advances the node to the next node if we can advance it
func (iter *DQIter[T]) Next() bool {
	if iter.IsValid() {
		iter.node = iter.node.next
		return true
	}
	return false
}

// Checks to see if we have a prev reference in the list
func (iter *DQIter[T]) HasPrev() bool {
	return iter.node != nil && iter.node.prev != nil
}

// Advances the iterator to the previous node if we can advance it
func (iter *DQIter[T]) Prev() bool {
	if iter.IsValid() {
		iter.node = iter.node.prev
		return true
	}
	return false
}

// Returns the payload of the current node that the iterator is pointing to
func (iter *DQIter[T]) GetPayload() T {
	return iter.node.payload
}

func (iter *DQIter[T]) IsValid() bool {
	return iter.node != nil
}
