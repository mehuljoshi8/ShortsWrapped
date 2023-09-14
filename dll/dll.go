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
	ll *LinkedList
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

// auxillary function to merge two sorted lists
// returns a merged list's head and tail
func merge(n1 *DLLNode, n2 *DLLNode, comp_fn LLPayloadComparatorFn) (*DLLNode, *DLLNode) {
    dummy := new(DLLNode)
    curr_dummy := dummy
    var selected, tmp *DLLNode
    for n1 != nil && n2 != nil {
        if comp_fn(n1.payload, n2.payload) > 0 {
            selected = n2
            n2 = n2.next
        } else {
            selected = n1
            n1 = n1.next
        }
        curr_dummy.next = selected
        selected.prev = curr_dummy
        tmp = selected.next
        selected.next = nil
        if tmp != nil {
            tmp.prev = nil
        }
        curr_dummy = curr_dummy.next
    }

    // if we have leftover
    if n1 != nil {
        curr_dummy.next = n1
        n1.prev = curr_dummy
    }
   
    if n2 != nil {
        curr_dummy.next = n2
        n2.prev = curr_dummy
    }
    
    // move curr_dummy to the end of the list
    for curr_dummy.next != nil {
        curr_dummy = curr_dummy.next    
    }

    return dummy.next, curr_dummy
}

// returns the head and tail of the newly merged list.
func mergeSort(n *DLLNode, comp_fn LLPayloadComparatorFn) (*DLLNode, *DLLNode) {
    var slow, fast *DLLNode
    slow = n
    fast = n.next
    if fast == nil {
        // case single elem list;
        return slow, slow
    }
    // find mid point using two pointers.
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
    }

    // now slow points to the middle elem.
    slow = slow.next
    slow.prev.next = nil
    slow.prev = nil
    // n is the start of the first list
    // slow is the start of the other one.
    left, _ := mergeSort(n, comp_fn)
    right, _ := mergeSort(slow, comp_fn)
    return merge(left, right, comp_fn)
}

// sorts the list that ll points to using merge sort.
func (ll *LinkedList) Sort(ascending bool, comparator_fn LLPayloadComparatorFn) {
    if ll.GetSize() == 0 {
        return
    }
    head, tail := mergeSort(ll.head, comparator_fn)
    ll.head = head
    ll.tail = tail
}

// for the given linked list at the pos (0 = head; 1 = tail)
func (ll *LinkedList) Iterator(pos int) *LLIter {
    ll_iter := new(LLIter)
    ll_iter.ll = ll
    if pos == 0 {
        ll_iter.node = ll.head
    } else {
        ll_iter.node = ll.tail
    }
    return ll_iter
}

// true if can advnace (not tail)
// false otherwise (tail)
func (ll_iter *LLIter) HasNext() bool {
    return ll_iter.node.next != nil
}

func (ll_iter *LLIter) Next() bool {
    if ll_iter.HasNext() {
        ll_iter.node = ll_iter.node.next
        return true
    }
    return false
}

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

// Returns the payload that the iterator is currently
// pointing at
func (ll_iter *LLIter) GetPayload() interface{} {
    return ll_iter.node.payload
}

// Delete the node the iterator is pointing to.  After deletion, the iterator:
//
// - invalid if the iterator was empty.
//
// - the successor of the deleted node, if there is one.
//
// - the predecessor of the deleted node, if the iterator was pointing at
//   the tail.
//
// Arguments:
//
// - iter:  the iterator to delete from
//
// - payload_free_function: invoked to free the payload
//
// Returns:
//
// - false if the deletion succeeded, but the list is now empty
//
// - true if the deletion succeeded, and the list is still non-empty
func (ll_iter *LLIter) Delete(payload_free_fn LLPayloadFreeFn) bool {
    if ll_iter.node == nil || ll_iter.ll == nil {
        // this is an invalid iterator
        return false
    }
    
    var p interface{}
    // handle single elem case (now list is empty) -> return false
    if ll_iter.ll.GetSize() == 1 {
        _, p = ll_iter.ll.Pop()
        payload_free_fn(p)
        ll_iter.node = nil
        return false
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

    payload_free_fn(p)
    return true
}

// TODO: Complete the InsertBefore function
func (ll_iter *LLIter) InsertBefore(p interface{}) bool {
    return false
}
