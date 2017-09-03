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

func TestDoNotInsertTheSameValue(t *testing.T) {
	tt := []struct {
		name            string
		insertValues    []int
		expectedResults []bool
	}{
		{"Insert root node", []int{1}, []bool{true}},
		{"Insert left node", []int{2, 1}, []bool{true, true}},
		{"Insert right node", []int{1, 2}, []bool{true, true}},
		{"Insert right and left node", []int{2, 1, 3}, []bool{true, true, true}},
		{
			"Insert with recursion node",
			[]int{5, 4, 1, 6, 9},
			[]bool{true, true, true, true, true},
		},
		{"Insert duplicate for root node", []int{1, 1}, []bool{true, false}},
		{
			"Insert children duplicate",
			[]int{2, 1, 1},
			[]bool{true, true, false},
		},
		{
			"Insert recursive duplicate",
			[]int{5, 4, 1, 8, 9, 1},
			[]bool{true, true, true, true, true, false},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			bst := trees.BST{}

			// Insert tree nodes
			for index, nodeValue := range tc.insertValues {
				result := bst.Insert(nodeValue)
				if result != tc.expectedResults[index] {
					t.Errorf("Expected result to be '%v'. Got '%v'", tc.expectedResults[index], result)
				}
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

			for i, value := range tc.insertValues[1:] {
				// if the insert result is true, the node should be allocated with the right value, otherwise
				// should be nil
				if tc.expectedResults[i+1] {
					checkValueInsert(t, bst.Root, value)
				} else {
					checkDuplicateValueInsert(t, bst.Root, value)
				}
			}
		})
	}
}

// If value is less than node value' then left node value'' or the children of left node should be value and vice versa
func checkValueInsert(t *testing.T, parentNode *trees.Node, value int) {
	if value < parentNode.Value {
		if parentNode.Left.Value != value {
			checkValueInsert(t, parentNode.Left, value)
		} else if parentNode == nil {
			t.Errorf("Expected left value to be '%v'. Got '%v'", value, parentNode.Left.Value)
		}
	} else {
		if parentNode.Right.Value != value {
			checkValueInsert(t, parentNode.Right, value)
		} else if parentNode == nil {
			t.Errorf("Expected left value to be '%v'. Got '%v'", value, parentNode.Left.Value)
		}
	}
}

// If the node value' is equal to the value, then their children values should be different from parent value or nil
func checkDuplicateValueInsert(t *testing.T, parentNode *trees.Node, value int) {
	if parentNode == nil {
		return
	}

	if parentNode.Value == value {
		if parentNode.Left != nil && parentNode.Left.Value == value {
			t.Errorf("Not expected duplicated value '%v' in left node", value)
		}

		if parentNode.Right != nil && parentNode.Right.Value == value {
			t.Errorf("Not expected duplicated value '%v' in right node", value)
		}
	} else {
		if value < parentNode.Value {
			checkDuplicateValueInsert(t, parentNode.Left, value)
		} else {
			checkDuplicateValueInsert(t, parentNode.Right, value)
		}
	}
}
