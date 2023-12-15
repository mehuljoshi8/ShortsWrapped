package dll

// Declaring a DLLNode with generics :)
type dLLNode[T interface{}] struct {
	next *dLLNode[T]
	prev *dLLNode[T]
	val  T
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

// push a value onto the front of a list
func (dll *Deque[T]) Push(value T) {
	var toInsert *dLLNode[T] = new(dLLNode[T])
	toInsert.val = value

	if dll.head == nil && dll.tail == nil {
		// case empty list
		dll.head = toInsert
		dll.tail = toInsert
	} else {
		// case non-empty list: just push it onto the head
		toInsert.next = dll.head
		dll.head.prev = toInsert
		dll.head = toInsert
	}

	dll.size++
}

// pops a value off the front of the list
func (dll *Deque[T]) Pop() (bool, T) {
	if dll.head == nil && dll.tail == nil {
		return false, zero[T]()
	}

	// let's think about what happens here
	// if it's a single list we do something different
	// if it's a non-singleton list we do soemthing different
	var value = dll.head.val
	if dll.head == dll.tail {
		// case singleton list.

	} else {

	}
	return true, value
}

// Return but not remove the first element.
func (dll *Deque[T]) Peek() (bool, T) {
	return false, zero[T]()
}

// push values onto the back of the list
func (dll *Deque[T]) Append(value T) {

}

// remove item from the end of the list
func (dll *Deque[T]) Slice() (bool, T) {
	return false, zero[T]()
}

// Return but not remove the last element.
func (dll *Deque[T]) PeekBack() (bool, T) {
	return false, zero[T]()
}

// get the size of the items
func (dll *Deque[T]) Size() int {
	return dll.size
}

// Sorts the Contents of the list
func (dll *Deque[T]) Sort(comp_fn comparatorFn) {

}
