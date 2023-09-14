package dll

import (
    "testing"
    "math/rand"
    "time"
)

func TestAllocateLinkedList(t *testing.T) {
    ll := AllocateLinkedList()
    got := ll.GetSize()
    var want uint64
    want = 0

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestPushPop(t *testing.T) {
    // we are going to test if push and pop work as expected
    // let's add nums 1,2,3,4,5,6 into the list
    ll := AllocateLinkedList()
    for i := 1; i <= 6; i++ {
        ll.Push(i)
        if ll.GetSize() != uint64(i) {
            t.Errorf("push not updating size: got %q, wanted %q", ll.GetSize(), i)
        }
    }

    for i := 6; i >= 1; i-- {
        _, v := ll.Pop()
        if v != i {
            t.Errorf("pop ordering incorrect: got %q, wanted %q", v, i)
        }

        if ll.GetSize() != uint64(i-1) {
            t.Errorf("pop not decrementing size: got %q, wanted %q", ll.GetSize(), i-1)
        }
    }

    b, v := ll.Pop()
    if b != false && v != nil {
        t.Errorf("pop on empty list not returning the correct values: got (%t, %v), wanted (%t, %v)", b, v, false, nil)
    }
}

func TestPushPopAppendSlice(t *testing.T) {
    ll := AllocateLinkedList()
    ll.Push(1)
    b, v :=  ll.Slice()
    if b != true && v != 1 {
        t.Errorf("Push Slice not working as expected")
    }
    
    for i := 0; i < 10; i++ {
        if i % 2 == 0 {
            // for every even number we are going to push a value onto the stack
            ll.Push(i)
        } else {
            // odd numbers get appended
            ll.Append(i)
        }
    }
    
    res := []int{8, 6, 4, 2, 0, 1, 3, 5, 7, 9}
    // then we are going to check the sequence from reverse by slicing all the elemnts off to make sure
    // they are the correct values.
    i := len(res) - 1
    for ll.GetSize() > 0 {
        _, v := ll.Slice()
        if v.(int) != res[i] {
            t.Errorf("append-slice-push not working as expected [expected: %q, actual: %q]", v.(int), res[i])
        }
        i--
    }

    ll.Append(11)
    _, v = ll.Pop()
    if v.(int) != 11 {
        t.Errorf("append to an empty list not working as expected")
    }

    b, v = ll.Slice()
    if b != false && v != nil {
        t.Errorf("slice on empty list returns non-empty")
    }
}

// Comparator function for dll
func comp_fn(p1 interface{}, p2 interface{}) int {
    return p1.(int) - p2.(int)
}

func TestSort(t *testing.T) {
    // create a linked list with random elems    
    ll := AllocateLinkedList()
    
    ll.Sort(true, comp_fn)
    
    if ll.GetSize() > 0 {
        t.Errorf("sorting on an empty list increased size")
    }

    ll.Push(1)
    ll.Sort(true, comp_fn)
    if ll.head.payload.(int) != 1 {
        t.Errorf("error sorting one element list failed")
    }
    ll.Pop()

    rand.Seed(time.Now().Unix())
    size := rand.Intn(10000)
    t.Log(size)
    for i := 0; i < size; i++ {
        ll.Push(rand.Intn(100))
    }

    ll.Sort(true, comp_fn)
    b, v := ll.Pop()
    count := 0
    if b {
        count = 1
    }
    for ll.GetSize() > 0 {
        _, v2 := ll.Pop()
        if v2.(int) < v.(int) {
            t.Errorf("error in sorting %v >= %v", v2.(int), v.(int))
        }
        v = v2
        count++
    }

    if count != size {
        t.Errorf("elements got lost after sorting: count = %v, size = %v", count, size)
    }
}

/*
func TestAllocateIterator(t *testing.T) {

}

func TestMoveIterator(t *testing.T) {

}

func TestIteratorGetPayload(t *testing.T) {

}

func TestIteratorDelete(t *testing.T) {

}

func TestIteratorInsertBefore(t *testing.T) {

}
*/
