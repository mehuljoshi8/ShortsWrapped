package dll

// A DLLNode contains a customer-supplied payload
// and next and prev pointers.
type DLLNode struct {
	payload *interface{}
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
type LLPayloadComparatorFn func(p1 *interface{}, p2 *interface{}) int
type LLPayloadFreeFn func(i *interface{}) *interface{}

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

// TODO: Complete Push
// adds a new element to the head of the linked list.
// arguments: the payload to push
func (ll *LinkedList) Push(payload *interface{}) {
	return
}

// TODO: Complete Pop
// pops an element from the head of the linked list
// and returns the payload that was provided nil otherwise
func (ll *LinkedList) Pop() (bool, *interface{}) {
	return false, nil
}

// TODO: Complete Append
// appends a new element to the tail of the list
// returns false on failure
func (ll *LinkedList) Append(payload *interface{}) {
	return
}

// TODO: Compelte Slice
// pops an element from the tail of the linked list.
// returns the payload and a boolean to the client.
func (ll *LinkedList) Slice() (bool, *interface{}) {
	return false, nil
}

// TODO: Write a sorting function to sort the linked list
func (ll *LinkedList) Sort(ascending bool, comparator_fn LLPayloadComparatorFn) {
	return
}

// TODO: Write the Iterator method which returns an LLIterator
// for the given linked list at the pos (0 = head; 1 = tail)
func (ll *LinkedList) Iterator(pos int) *LLIter {
    return nil
}

// TODO: Complete the implementation of has next.
// true if can advnace (not tail)
// false otherwise (tail)
func (ll_iter *LLIter) HasNext() bool {
    return false
}

// TODO: Complete Iteration of Next
func (ll_iter *LLIter) Next() bool {
    return false
}

// TODO: Complete HasPrev implementation
func (ll_iter *LLIter) HasPrev() bool {
    return false
}

// TODO: complete getPayload
// Returns the payload that the iterator is currently
// pointing at
func (ll_iter *LLIter) GetPayload() *interface{} {
    return nil
}

// TODO: Complete the delete function
func (ll_iter *LLIter) Delete(payload_free_fn LLPayloadFreeFn) bool {
    return false
}

// TODO: Complete the InsertBefore function
func (ll_iter *LLIter) InsertBefore(p *interface{}) bool {
    return false
}
