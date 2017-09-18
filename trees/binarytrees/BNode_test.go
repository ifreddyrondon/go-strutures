package binarytrees_test

import (
	"testing"

	"github.com/ifreddyrondon/go-strutures/trees/binarytrees"
)

func TestNewNode(t *testing.T) {
	var value int
	node := binarytrees.NewNode(value)

	if node.Value != value {
		t.Errorf("Expected NewNode value to be '%v'. Got '%v'", value, node.Value)
	}

	if node.Left != nil {
		t.Errorf("Expected NewNode Left to be nil. Got '%v'", node.Left)
	}

	if node.Right != nil {
		t.Errorf("Expected NewNode Left to be nil. Got '%v'", node.Left)
	}
}
