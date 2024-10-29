package graph

import (
	"github.com/sosalejandro/algo-practice/data-structures/common"
	simple_queue "github.com/sosalejandro/algo-practice/data-structures/simple-queue"
	simple_stack "github.com/sosalejandro/algo-practice/data-structures/simple-stack"
)

// TraversalStrategy defines the type of traversal strategy.
type TraversalStrategy int

const (
	StackTraversal TraversalStrategy = iota
	QueueTraversal
)

// Transporter is a generic interface for traversal methods, supporting Stack and Queue
// so that HasPath can use either data structure.
type Transporter[T any] interface {
	// Next returns the next item in the transporter.
	Next() (common.Item[T], error)
	// Add adds an item to the transporter.
	Add(element T) error
	// IsEmpty checks if the transporter is empty.
	IsEmpty() bool
}

// StackTransporter adapts a stack to implement the generic Transporter interface.
type StackTransporter[T any] struct {
	stack       *simple_stack.Stack[T, common.Item[T]]
	itemFactory common.ItemFactory[T]
}

// NewStackTransporter creates a new generic StackTransporter.
func NewStackTransporter[T any](itemFactory common.ItemFactory[T]) *StackTransporter[T] {
	return &StackTransporter[T]{simple_stack.NewStack[T, common.Item[T]](), itemFactory}
}

func (s *StackTransporter[T]) Next() (common.Item[T], error) {
	return s.stack.Pop()
}

func (s *StackTransporter[T]) Add(element T) error {
	return s.stack.Push(s.itemFactory(element))
}

func (s *StackTransporter[T]) IsEmpty() bool {
	return s.stack.IsEmpty()
}

// QueueTransporter adapts a queue to implement the generic Transporter interface.
type QueueTransporter[T any] struct {
	queue       *simple_queue.Queue[T, common.Item[T]]
	itemFactory common.ItemFactory[T]
}

// NewQueueTransporter creates a new generic QueueTransporter.
func NewQueueTransporter[T any](itemFactory common.ItemFactory[T]) *QueueTransporter[T] {
	return &QueueTransporter[T]{simple_queue.NewQueue[T, common.Item[T]](), itemFactory}
}

func (q *QueueTransporter[T]) Next() (common.Item[T], error) {
	return q.queue.Dequeue()
}

func (q *QueueTransporter[T]) Add(element T) error {
	return q.queue.Enqueue(q.itemFactory(element))
}

func (q *QueueTransporter[T]) IsEmpty() bool {
	return q.queue.IsEmpty()
}

// NewTransporter creates a new generic Transporter based on the traversal strategy.
func NewTransporter[T any](strategy TraversalStrategy, itemFactory common.ItemFactory[T]) Transporter[T] {
	switch strategy {
	case StackTraversal:
		return NewStackTransporter(itemFactory)
	case QueueTraversal:
		return NewQueueTransporter(itemFactory)
	default:
		return nil
	}
}
