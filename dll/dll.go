package dll

// A DLLNode contains a customer-supplied payload
// and next and prev pointers.
type DLLNode struct {
	payload interface{}
	next    *DLLNode
	prev    *DLLNode
}

// The linked list struct manages the Nodes in the list.
type LinkedList struct {
	size uint64
	head *DLLNode
	tail *DLLNode
}

// This struct represents the state of an iterator.
type LLIter struct {
	list LinkedList
	node *DLLNode
}

// A type alias for a comparator function that the client has to define
type LLPayloadComparatorFn func(p1 interface{}, p2 interface{}) int
type LLPayloadFreeFn func(i interface{}) interface{}

// allocate a new linked list and returns it to client
func AllocateLinkedList() *LinkedList {
	//ll = new LinkedList
	var ll *LinkedList
	ll = new(LinkedList)
	ll.size = 0
	ll.head = nil
	ll.tail = nil
	return ll
}

// Returns the number of elements
func (ll *LinkedList) GetSize() uint64 {
	return ll.size
}

// adds a new element to the head of the linked list.
// arguments: the payload to push
func (ll *LinkedList) Push(payload interface{}) {
    // have to handle the case where there are no elems
    node := new(DLLNode)
    node.payload = payload
    node.next = ll.head
    node.prev = nil
    if ll.head != nil {
        ll.head.prev = node
    } else {
        ll.tail = node
    }
    ll.head = node
    ll.size++
}

// pops an element from the head of the linked list
// and returns the payload that was provided nil otherwise
func (ll *LinkedList) Pop() (bool, interface{}) {
	if ll.head != nil {
        node := ll.head
        ll.head = ll.head.next
        node.next = nil
        if ll.head != nil {
            ll.head.prev = nil
        } else {
            // in the 1 node case
            ll.tail = nil
        }
        ll.size--
        return true, node.payload
    }
    return false, nil
}

// appends a new element to the tail of the list
// returns false on failure
func (ll *LinkedList) Append(payload interface{}) {
    node := new(DLLNode)
    node.payload = payload
    node.prev = ll.tail 
    node.next = nil
    if ll.tail != nil {
        ll.tail.next = node
    } else {
        ll.head = node
    }
    ll.tail = node
    ll.size++
}

// pops an element from the tail of the linked list.
// returns the payload and a boolean to the client.
func (ll *LinkedList) Slice() (bool, interface{}) {
    if ll.tail != nil {
        node := ll.tail
        ll.tail = ll.tail.prev
        node.prev = nil
        if ll.tail != nil {
            ll.tail.next = nil
        } else {
            ll.head = nil
        }
        ll.size--
        return true, node.payload
    }
    return false, nil
}

// TODO: Write a sorting function to sort the linked list
func (ll *LinkedList) Sort(ascending bool, comparator_fn LLPayloadComparatorFn) {
	// mergesort on a linked list is kinda cool
     
}

// TODO: Write the Iterator method which returns an LLIterator
// for the given linked list at the pos (0 = head; 1 = tail)
func (ll *LinkedList) Iterator(pos int) *LLIter {
    ll_iter := new(LLIter)
    ll_iter.list = ll
    if pos == 0 {
        ll_iter.node = ll.head
    } else {
        ll_iter.node = ll.tail
    }
    return ll_iter
}

// TODO: Complete the implementation of has next.
// true if can advnace (not tail)
// false otherwise (tail)
func (ll_iter *LLIter) HasNext() bool {
    return ll_iter.node.next != nil
}

// TODO: Complete Iteration of Next
func (ll_iter *LLIter) Next() bool {
    if ll_iter.HasNext() {
        ll_iter.node = ll_iter.node.next
        return true
    }
    return false
}

// TODO: Complete HasPrev implementation
func (ll_iter *LLIter) HasPrev() bool {
    return ll_iter.node.prev != nil
}

func (ll_iter *LLIter) Prev() bool {
    if ll_iter.HasPrev() {
        ll_iter.node = ll_iter.node.prev
        return true
    }
    return false
}

// TODO: complete getPayload
// Returns the payload that the iterator is currently
// pointing at
func (ll_iter *LLIter) GetPayload() interface{} {
    return ll_iter.node.payload
}

// TODO: Complete the delete function
func (ll_iter *LLIter) Delete(payload_free_fn LLPayloadFreeFn) bool {
    return false
}

// TODO: Complete the InsertBefore function
func (ll_iter *LLIter) InsertBefore(p interface{}) bool {
    return false
}
