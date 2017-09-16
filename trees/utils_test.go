package trees_test

import (
	"testing"

	"github.com/ifreddyrondon/go-strutures/trees"
)

func TestNewRandBST(t *testing.T) {
	bst := trees.NewRandBST(10)

	if bst.Root == nil {
		t.Error("Expected root to be not nil")
	}

	if bst.Len() != 10 {
		t.Errorf("Expected Len value to be '%v'. Got '%v'", 10, bst.Len())
	}
}
