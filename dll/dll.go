package dll

// LLPayload type definition:
// For generality, a payload must be large enough to hold a pointer.
// If the client's data is no bigger than a pointer, a copy of that
// data can be stored in the LinkedList, by casting it to the LLPayload
// type.  Otherwise, a pointer to the client's data is maintained in
// the list.
type LLPayload_t = interface{}

// A DLLNode contains a customer-supplied payload
// and next and prev pointers.
type DLLNode struct {
	payload LLPayload_t
	next    *DLLNode
	prev    *DLLNode
}

// The linked list struct manages the Nodes in the list.
type LinkedList struct {
	size uint64
	head *DLLNode
	tail *DLLNode
}

// A type alias for a comparator function that the client has to define
type LLPayloadComparatorFn func(p1 LLPayload_t, p2 LLPayload_t) int

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
func (ll *LinkedList) Push(payload LLPayload_t) {
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
func (ll *LinkedList) Pop() (bool, LLPayload_t) {
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
func (ll *LinkedList) Append(payload LLPayload_t) {
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
func (ll *LinkedList) Slice() (bool, LLPayload_t) {
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
func (ll *LinkedList) Sort(comparator_fn LLPayloadComparatorFn) {
    if ll.GetSize() == 0 {
        return
    }
    head, tail := mergeSort(ll.head, comparator_fn)
    ll.head = head
    ll.tail = tail
}
