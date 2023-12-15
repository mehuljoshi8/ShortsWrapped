package dll

// Declaring a DLLNode with generics :)
type dLLNode[T interface{}] struct {
	next    *dLLNode[T]
	prev    *dLLNode[T]
	payload T
}

// Deque struct for a doubly linked list
type Deque[T interface{}] struct {
	head *dLLNode[T]
	tail *dLLNode[T]
	size int
}

// Returns an empty parameter type T variable upon construction
// Used throughout the DLL package
func zero[T interface{}]() T {
	var zeroVal T
	return zeroVal
}

// Comparator function type for functions to abide to they just need to compare two
// like things and return some kind of integer and then we say that it's a comparator function
// cause that is really what it is at the end of hte day.
type comparatorFn func(a, b interface{}) int

// Constructor for a DoublyLinkedList
func NewDeque[T interface{}]() Deque[T] {
	return Deque[T]{}
}

// Push a value onto the front of a list
func (dll *Deque[T]) Push(payload T) {
	var toInsert *dLLNode[T] = new(dLLNode[T])
	toInsert.payload = payload

	if dll.head == nil && dll.tail == nil {
		dll.head = toInsert
		dll.tail = toInsert
	} else {
		toInsert.next = dll.head
		dll.head.prev = toInsert
		dll.head = toInsert
	}

	dll.size++
}

// Pops a value off the front of the list
func (dll *Deque[T]) Pop() (bool, T) {
	if dll.head == nil && dll.tail == nil {
		return false, zero[T]()
	}

	var payload = dll.head.payload
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
	}
	dll.size--
	return true, payload
}

// Return but not remove the first element.
func (dll *Deque[T]) Peek() (bool, T) {
	if dll.head == nil && dll.tail == nil {
		return false, zero[T]()
	}
	return true, dll.head.payload
}

// push values onto the back of the list
func (dll *Deque[T]) Append(payload T) {
	var toInsert *dLLNode[T] = new(dLLNode[T])
	toInsert.payload = payload

	if dll.head == nil && dll.tail == nil {
		dll.head = toInsert
		dll.tail = toInsert
	} else {
		toInsert.prev = dll.tail
		dll.tail.next = toInsert
		dll.tail = toInsert
	}
	dll.size++
}

// remove item from the end of the list
func (dll *Deque[T]) Slice() (bool, T) {
	if dll.head == nil && dll.tail == nil {
		return false, zero[T]()
	}

	var value = dll.tail.payload
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
	}
	dll.size--
	return true, value
}

// Return but not remove the last element.
func (dll *Deque[T]) PeekBack() (bool, T) {
	if dll.head == nil && dll.tail == nil {
		return false, zero[T]()
	}
	return true, dll.tail.payload
}

// get the size of the items
func (dll *Deque[T]) Size() int {
	return dll.size
}

// Auxillary function to merge two doubly linked lists
func merge[T interface{}](n1 *dLLNode[T], n2 *dLLNode[T], comp_fn comparatorFn) (*dLLNode[T], *dLLNode[T]) {
	dummy := new(dLLNode[T])
	curr_dummy := dummy
	var selected, tmp *dLLNode[T]
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
func mergeSort[T interface{}](n *dLLNode[T], comp_fn comparatorFn) (*dLLNode[T], *dLLNode[T]) {
	var slow, fast *dLLNode[T]
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

// Sorts the Contents of the list
func (dll *Deque[T]) Sort(comp_fn comparatorFn) {
	if dll.size == 0 {
		return
	}

	head, tail := mergeSort(dll.head, comp_fn)
	dll.head = head
	dll.tail = tail
}
