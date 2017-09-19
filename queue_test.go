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

