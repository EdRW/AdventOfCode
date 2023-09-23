package utils

import (
	"errors"
	"fmt"
)

// Queue is a channel with FIFO added operations
type Queue[T comparable] chan T

func NewQueue[T comparable](capacity int) Queue[T] {
	return make(Queue[T], capacity)
}

// Add a value to the end of the queue
// and block until queue capacity is available
func (ch Queue[T]) BlockingEnqueue(value T) {
	ch <- value
}

// Add a value to the end of the queue
// and return an error if queue capacity
// is would be exceeded
func (ch Queue[T]) Enqueue(value T) error {
	if len(ch) >= cap(ch) {
		return fmt.Errorf("Queue: %v capacity (%d) exceeded, could not enqueue value %v", &ch, cap(ch), value)
	}
	ch.BlockingEnqueue(value)
	return nil
}

// Remove and return the value at the front of the queue
// and block until a value is available
func (q Queue[T]) BlockingDequeue() T {
	return <-q
}

// Remove and return the value at the front of the queue
// and return an error if the queue is empty
func (ch Queue[T]) Dequeue() (T, error) {

	if len(ch) == 0 {
		var value T
		return value, errors.New("Dequeue called on empty queue")
	}

	return ch.BlockingDequeue(), nil
}
