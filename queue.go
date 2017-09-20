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


// Peek returns but does not remove, the head of this queue, or returns nil if this queue is empty.
func (q *Queue) Peek() Item {
	var item Item
	if q.Size() > 0 {
		item = (*q)[0]
	}
	return item
}

// Size returns the number of Items in the queue
func (q *Queue) Size() int {
	return len(*q)
}
