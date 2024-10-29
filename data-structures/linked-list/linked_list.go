package linkedlist

type NodeValue interface {
	string | int | uint | float32 | float64 | int64 | uint64
}

type NodeArrayValue interface {
	[]string | []int | []uint | []float32 | []float64 | []int64 | []uint64
}

type Node[T NodeValue | NodeArrayValue] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T NodeValue | NodeArrayValue] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (ll *LinkedList[T]) First() *Node[T] {
	return ll.Head
}

func (ll *LinkedList[T]) Last() *Node[T] {
	return ll.Tail
}

func (ll *LinkedList[T]) AddFirst(value T) {
	newNode := &Node[T]{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = newNode
	}
}

func (ll *LinkedList[T]) AddLast(value T) {
	newNode := &Node[T]{Value: value}
	if ll.Tail == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
}

func (ll *LinkedList[T]) RemoveFirst() *Node[T] {
	if ll.Head == nil {
		return nil
	}
	removedNode := ll.Head
	ll.Head = ll.Head.Next
	if ll.Head == nil {
		ll.Tail = nil
	}
	return removedNode
}

func (ll *LinkedList[T]) RemoveLast() *Node[T] {
	if ll.Head == nil {
		return nil
	}
	if ll.Head == ll.Tail {
		removedNode := ll.Head
		ll.Head = nil
		ll.Tail = nil
		return removedNode
	}
	current := ll.Head
	for current.Next != ll.Tail {
		current = current.Next
	}
	removedNode := ll.Tail
	ll.Tail = current
	ll.Tail.Next = nil
	return removedNode
}

func (ll *LinkedList[T]) IndexAt(index int) *Node[T] {
	current := ll.Head
	for i := 0; current != nil && i < index; i++ {
		current = current.Next
	}
	return current
}

func (ll *LinkedList[T]) RemoveAt(index int) *Node[T] {
	if index == 0 {
		return ll.RemoveFirst()
	}
	prev := ll.IndexAt(index - 1)
	if prev == nil || prev.Next == nil {
		return nil
	}
	removedNode := prev.Next
	prev.Next = removedNode.Next
	if removedNode == ll.Tail {
		ll.Tail = prev
	}
	return removedNode
}

func (ll *LinkedList[T]) AddAt(index int, value T) {
	if index == 0 {
		ll.AddFirst(value)
		return
	}
	prev := ll.IndexAt(index - 1)
	if prev == nil {
		return
	}
	newNode := &Node[T]{Value: value, Next: prev.Next}
	prev.Next = newNode
	if newNode.Next == nil {
		ll.Tail = newNode
	}
}
