package gostrutures

type Queue []Item

// New build a Queue with the root.
func New() *Queue {
	return &Queue{}
}

// Push adds an element (Item) to the end of the queue.
func (q *Queue) Push(value Item) {
	*q = append(*q, value)
}

