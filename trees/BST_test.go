package trees_test

import (
	"testing"

	"github.com/ifreddyrondon/go-strutures/trees"
)

func TestNewConstructor(t *testing.T) {
	bst := trees.New(1)
	if bst.Root == nil {
		t.Error("Expected root to be not nil")
	}

	if bst.Root.Value != 1 {
		t.Errorf("Expected root value to be '1'. Got '%v'", bst.Root.Value)
	}
}

func TestInsert(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
	}{
		{"Insert one node", []int{1}},
		{"Insert left node", []int{2, 1}},
		{"Insert right node", []int{1, 2}},
		{"Insert right and left node", []int{2, 1, 3}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}

			// Insert tree nodes
			for _, nodeValue := range tc.insertValues {
				bst.Insert(nodeValue)
			}

			if bst.Root == nil {
				t.Error("Expected root to be not nil")
			}

			if bst.Root.Value != tc.insertValues[0] {
				t.Errorf("Expected root value to be '%v'. Got '%v'", tc.insertValues[0], bst.Root.Value)
			}

			if len(tc.insertValues) == 1 {
				return
			}

			for _, value := range tc.insertValues[1:] {
				if value < bst.Root.Value {
					if bst.Root.Left.Value != value {
						t.Errorf("Expected left value to be '%v'. Got '%v'", bst.Root.Left.Value, value)
					}
				} else {
					if bst.Root.Right.Value != value {
						t.Errorf("Expected right value to be '%v'. Got '%v'", bst.Root.Right.Value, value)
					}
				}
			}
		})
	}

}
