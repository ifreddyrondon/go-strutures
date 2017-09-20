package gostrutures_test

import (
	"testing"

	"github.com/ifreddyrondon/gostrutures"
)

func TestNew(t *testing.T) {
	q := gostrutures.New()
	if q == nil {
		t.Error("Expected queue to be not nil")
	}
}

func TestQueue_Push(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
	}{
		{"empty queue", []int{}},
		{"queue with elements", []int{3, 2, 3, 4}},
	}

	for _, tc := range tt {
		queue := new(gostrutures.Queue)
		// Insert tree nodes
		for _, nodeValue := range tc.insertValues {
			queue.Push(nodeValue)
		}

		if queue == nil {
			t.Error("Expected queue to be not nil")
		}

		for i := range tc.insertValues {
			if (*queue)[i] != tc.insertValues[i] {
				t.Errorf("Expected items into the queue to be '%v'. Got '%v'", tc.insertValues[i], (*queue)[i])
				break
			}
		}
	}
}

func TestQueue_Size(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
	}{
		{"empty queue", []int{}},
		{"queue with elements", []int{3, 2, 3, 4}},
	}

	for _, tc := range tt {
		queue := new(gostrutures.Queue)
		// Insert tree nodes
		for _, nodeValue := range tc.insertValues {
			queue.Push(nodeValue)
		}

		if queue.Size() != len(tc.insertValues) {
			t.Errorf("Expected queue size to be '%v'. Got '%v'", len(tc.insertValues), queue.Size())
		}
	}
}

func TestQueue_Peek(t *testing.T) {
	// GIVEN
	queue := new(gostrutures.Queue)
	insertValues, popElement := []int{3, 2, 5, 4}, 3
	// Insert tree nodes
	for _, nodeValue := range insertValues {
		queue.Push(nodeValue)
	}
	// When peek
	result := queue.Peek()
	// Then
	if result != popElement {
		t.Errorf("Expected pop element to be '%v'. Got '%v'", popElement, result)
	}

	if queue.Size() != len(insertValues) {
		t.Errorf("Expected queue size after pop to be '%v'. Got '%v'", len(insertValues), queue.Size())
	}
}

func TestQueue_PeekFromEmptyQueue(t *testing.T) {
	// GIVEN
	queue := new(gostrutures.Queue)
	// When peek
	result := queue.Peek()
	//Then
	if result != nil {
		t.Errorf("Expected pop element to be nil. Got '%v'", result)
	}

	if queue.Size() != 0 {
		t.Errorf("Expected queue size after pop to be '%v'. Got '%v'", 0, queue.Size())
	}
}

func TestQueue_Pop(t *testing.T) {
	// Given
	queue := new(gostrutures.Queue)
	popElement, lenAfterPop := 3, 3
	// Insert tree nodes
	for _, nodeValue := range []int{3, 2, 5, 4} {
		queue.Push(nodeValue)
	}
	// When pop
	result := queue.Pop()
	// Then
	if result != popElement {
		t.Errorf("Expected pop element to be '%v'. Got '%v'", popElement, result)
	}

	if queue.Size() != lenAfterPop {
		t.Errorf("Expected queue size after pop to be '%v'. Got '%v'", lenAfterPop, queue.Size())
	}
}

func TestQueue_PopFromEmptyQueue(t *testing.T) {
	// Given
	queue := new(gostrutures.Queue)
	// When pop
	result := queue.Pop()
	// Then
	if result != nil {
		t.Errorf("Expected pop element to be nil. Got '%v'", result)
	}

	if queue.Size() != 0 {
		t.Errorf("Expected queue size after pop to be '%v'. Got '%v'", 0, queue.Size())
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		expected     bool
	}{
		{"empty queue", []int{}, true},
		{"queue with elements", []int{3, 2, 3, 4}, false},
	}

	for _, tc := range tt {
		queue := new(gostrutures.Queue)
		// Insert tree nodes
		for _, nodeValue := range tc.insertValues {
			queue.Push(nodeValue)
		}

		if queue.IsEmpty() != tc.expected {
			t.Errorf("Expected queue IsEmpty to be '%v'. Got '%v'", tc.expected, queue.IsEmpty())
		}
	}
}
