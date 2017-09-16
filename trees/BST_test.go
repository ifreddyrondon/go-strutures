package trees_test

import (
	"testing"

	"bytes"

	"fmt"

	"github.com/ifreddyrondon/go-strutures/trees"
)

func TestNewConstructor(t *testing.T) {
	bst := trees.New(1)
	if bst.Root() == nil {
		t.Error("Expected root to be not nil")
	}

	if bst.Root().Value != 1 {
		t.Errorf("Expected root value to be '1'. Got '%v'", bst.Root().Value)
	}
}

func TestInsert(t *testing.T) {
	tt := []struct {
		name            string
		insertValues    []int
		expectedResults []bool
		expectedLen     int
	}{
		{"insert root node", []int{1}, []bool{true}, 1},
		{"insert left node (plain tree)", []int{2, 1}, []bool{true, true}, 2},
		{"insert right node (plain tree)", []int{1, 2}, []bool{true, true}, 2},
		{
			"insert right and left node (plain tree)",
			[]int{2, 1, 3},
			[]bool{true, true, true},
			3,
		},
		{
			"insert with recursion node",
			[]int{5, 4, 1, 6, 9},
			[]bool{true, true, true, true, true},
			5,
		},
		{"insert duplicate for root node", []int{1, 1}, []bool{true, false}, 1},
		{
			"insert children duplicate",
			[]int{2, 1, 1},
			[]bool{true, true, false},
			2,
		},
		{
			"insert recursive duplicate",
			[]int{5, 4, 1, 8, 9, 1},
			[]bool{true, true, true, true, true, false},
			5,
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

			if bst.Root().Value != tc.insertValues[0] {
				t.Errorf("Expected root value to be '%v'. Got '%v'", tc.insertValues[0], bst.Root().Value)
			}

			if bst.Len() != tc.expectedLen {
				t.Errorf("Expected Len value to be '%v'. Got '%v'", tc.expectedLen, bst.Len())
			}

			if len(tc.insertValues) == 1 {
				return
			}

			for i, value := range tc.insertValues[1:] {
				// if the insert result is true, the node should be allocated with the right value, otherwise
				// should be nil
				if tc.expectedResults[i+1] {
					checkValueInsert(t, bst.Root(), value)
				} else {
					checkDuplicateValueInsert(t, bst.Root(), value)
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

func fillTreeWithList(bst *trees.BST, list []int) {
	for _, v := range list {
		bst.Insert(v)
	}
}

func TestInOrderTraverse(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		expected     []int
	}{
		{"balanced tree", []int{5, 3, 1, 4, 7, 9, 6}, []int{1, 3, 4, 5, 6, 7, 9}},
		{"duplicate values", []int{5, 3, 1, 1, 7, 9, 9}, []int{1, 3, 5, 7, 9}},
		{"bst unbalanced to right", []int{5, 6, 7, 8, 9, 10}, []int{5, 6, 7, 8, 9, 10}},
		{"bst unbalanced to left", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"only root", []int{5}, []int{5}},
		{"nil root", []int{}, []int{}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			var result []int
			bst.InOrderTraverse(func(i int) {
				result = append(result, i)
			})
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("Expected in order traversal to be '%v'. Got '%v'", tc.expected, result)
					break
				}
			}
		})
	}
}

func TestPreOrderTraverse(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		expected     []int
	}{
		{"balanced tree", []int{5, 3, 1, 4, 7, 9, 6}, []int{5, 3, 1, 4, 7, 6, 9}},
		{"duplicate values", []int{5, 3, 1, 1, 7, 9, 9}, []int{5, 3, 1, 7, 9}},
		{"bst unbalanced to right", []int{5, 6, 7, 8, 9, 10}, []int{5, 6, 7, 8, 9, 10}},
		{"bst unbalanced to left", []int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{"only root", []int{5}, []int{5}},
		{"nil root", []int{}, []int{}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			var result []int
			bst.PreOrderTraverse(func(i int) {
				result = append(result, i)
			})
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("Expected pre order traversal to be '%v'. Got '%v'", tc.expected, result)
					break
				}
			}
		})
	}
}

func TestPostOrderTraverse(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		expected     []int
	}{
		{"balanced tree", []int{5, 3, 1, 4, 7, 9, 6}, []int{1, 4, 3, 6, 9, 7, 5}},
		{"duplicate values", []int{5, 3, 1, 1, 7, 9, 9}, []int{1, 3, 9, 7, 5}},
		{"bst unbalanced to right", []int{5, 6, 7, 8, 9, 10}, []int{10, 9, 8, 7, 6, 5}},
		{"bst unbalanced to left", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"only root", []int{5}, []int{5}},
		{"nil root", []int{}, []int{}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			var result []int
			bst.PostOrderTraverse(func(i int) {
				result = append(result, i)
			})
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("Expected post order traversal to be '%v'. Got '%v'", tc.expected, result)
					break
				}
			}
		})
	}
}

func TestGetMin(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		expected     *trees.Node
	}{
		{"balanced tree", []int{5, 3, 1, 4, 7, 9, 6}, &trees.Node{Value: 1, Left: nil, Right: nil}},
		{"duplicate values", []int{5, 3, 1, 1, 7, 9, 9}, &trees.Node{Value: 1, Left: nil, Right: nil}},
		{"bst unbalanced to right", []int{5, 6, 7, 8, 9, 10}, &trees.Node{Value: 5, Left: nil, Right: nil}},
		{"bst unbalanced to left", []int{5, 4, 3, 2, 1}, &trees.Node{Value: 1, Left: nil, Right: nil}},
		{"only root", []int{5}, &trees.Node{Value: 5, Left: nil, Right: nil}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			result := bst.Min()
			if result.Value != tc.expected.Value {
				t.Errorf("Expected min to be '%v'. Got '%v'", tc.expected, result)
			}
		})
	}
}

func TestGetMinForNilBST(t *testing.T) {
	bst := trees.BST{}
	result := bst.Min()
	if result != nil {
		t.Errorf("Expected min to be '%v'. Got '%v'", nil, result)
	}
}

func TestGetMax(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		expected     *trees.Node
	}{
		{"balanced tree", []int{5, 3, 1, 4, 7, 9, 6}, &trees.Node{Value: 9, Left: nil, Right: nil}},
		{"duplicate values", []int{5, 3, 1, 1, 7, 9, 9}, &trees.Node{Value: 9, Left: nil, Right: nil}},
		{"bst unbalanced to right", []int{5, 6, 7, 8, 9, 10}, &trees.Node{Value: 10, Left: nil, Right: nil}},
		{"bst unbalanced to left", []int{5, 4, 3, 2, 1}, &trees.Node{Value: 5, Left: nil, Right: nil}},
		{"only root", []int{5}, &trees.Node{Value: 5, Left: nil, Right: nil}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			result := bst.Max()
			if result.Value != tc.expected.Value {
				t.Errorf("Expected max to be '%v'. Got '%v'", tc.expected, result)
			}
		})
	}
}

func TestGetMaxForNilBST(t *testing.T) {
	bst := trees.BST{}
	result := bst.Max()
	if result != nil {
		t.Errorf("Expected max to be '%v'. Got '%v'", nil, result)
	}
}

func TestSearchNode(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		searchValue  int
		expected     *trees.Node
	}{
		{"balanced tree", []int{5, 3, 1, 4, 7, 9, 6}, 4, &trees.Node{Value: 4, Left: nil, Right: nil}},
		{"search duplicate values", []int{5, 3, 1, 1, 7, 9, 9}, 1, &trees.Node{Value: 1, Left: nil, Right: nil}},
		{"bst unbalanced to right", []int{5, 6, 7, 8, 9, 10}, 10, &trees.Node{Value: 10, Left: nil, Right: nil}},
		{"bst unbalanced to left", []int{5, 4, 3, 2, 1}, 1, &trees.Node{Value: 1, Left: nil, Right: nil}},
		{"only root", []int{5}, 5, &trees.Node{Value: 5, Left: nil, Right: nil}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			result := bst.Search(tc.searchValue)
			if result.Value != tc.expected.Value {
				t.Errorf("Expected search to be '%v'. Got '%v'", tc.expected, result)
			}
		})
	}
}

func TestSearchNilReturn(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		searchValue  int
	}{
		{"not found", []int{5, 3, 1, 4, 7, 9, 6}, 2},
		{"nil tree", []int{}, 1},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			result := bst.Search(tc.searchValue)
			if result != nil {
				t.Errorf("Expected search to be nil. Got '%v'", result)
			}
		})
	}
}

func TestHasNode(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		searchValue  int
		expected     bool
	}{
		{"balanced tree", []int{5, 3, 1, 4, 7, 9, 6}, 4, true},
		{"search duplicate values", []int{5, 3, 1, 1, 7, 9, 9}, 1, true},
		{"bst unbalanced to right", []int{5, 6, 7, 8, 9, 10}, 10, true},
		{"bst unbalanced to left", []int{5, 4, 3, 2, 1}, 1, true},
		{"only root", []int{5}, 5, true},
		{"not found", []int{5, 3, 1, 4, 7, 9, 6}, 2, false},
		{"nil tree", []int{}, 1, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			result := bst.Has(tc.searchValue)
			if result != tc.expected {
				t.Errorf("Expected has to be '%v'. Got '%v'", tc.expected, result)
			}
		})
	}
}

func TestRemoveNode(t *testing.T) {
	tt := []struct {
		name            string
		insertValues    []int
		deleteValue     int
		inOrderExpected []int
		resultExpected  bool
	}{
		// LEAF NODES
		{
			"remove root when len 1",
			[]int{5},
			5,
			[]int{},
			true,
		},
		{
			"remove leaf node from left branch when len 3 (plain)",
			[]int{2, 1, 3},
			1,
			[]int{2, 3},
			true,
		},
		{
			"remove leaf node from right branch when len 3 (plain)",
			[]int{2, 1, 3},
			3,
			[]int{1, 2},
			true,
		},
		{
			"remove leaf node from left branch",
			[]int{5, 7, 2, 1, 4, 6},
			1,
			[]int{2, 4, 5, 6, 7},
			true,
		},
		{
			"remove leaf node from right branch",
			[]int{5, 7, 2, 1, 4, 6},
			1,
			[]int{2, 4, 5, 6, 7},
			true,
		},
		{
			"remove leaf node when tree unbalanced to left",
			[]int{5, 4, 3, 2, 1},
			1,
			[]int{2, 3, 4, 5},
			true,
		},
		{
			"remove leaf node when tree unbalanced to right",
			[]int{5, 6, 7, 8, 9, 10},
			10,
			[]int{5, 6, 7, 8, 9},
			true,
		},
		// HALF-LEAF
		{
			"remove half-leaf from linear left branch",
			[]int{3, 2, 1, 4, 6},
			2,
			[]int{1, 3, 4, 6},
			true,
		},
		{
			"remove half-leaf from not linear left branch",
			[]int{5, 2, 3, 6},
			2,
			[]int{3, 5, 6},
			true,
		},
		{
			"remove half-leaf from linear right branch",
			[]int{3, 2, 1, 4, 6},
			4,
			[]int{1, 2, 3, 6},
			true,
		},
		{
			"remove half-leaf from not linear right branch",
			[]int{5, 2, 8, 7},
			8,
			[]int{2, 5, 7},
			true,
		},
		{
			"remove root node when tree unbalanced to left",
			[]int{5, 4, 3, 2, 1},
			5,
			[]int{1, 2, 3, 4},
			true,
		},
		{
			"remove leaf node when tree unbalanced to right",
			[]int{5, 6, 7, 8, 9, 10},
			5,
			[]int{6, 7, 8, 9, 10},
			true,
		},
		// INNER NODE
		{
			"remove root node when len 3 (plain)",
			[]int{2, 1, 3},
			2,
			[]int{1, 3},
			true,
		},
		{
			"remove root node",
			[]int{5, 7, 2, 1, 4, 6, 8},
			5,
			[]int{1, 2, 4, 6, 7, 8},
			true,
		},
		{
			"inner node from left branch",
			[]int{5, 7, 2, 1, 4, 6},
			2,
			[]int{1, 4, 5, 6, 7},
			true,
		},
		{
			"inner node from right branch",
			[]int{5, 2, 7, 6, 8},
			7,
			[]int{2, 5, 6, 8},
			true,
		},
		// NOT FOUND
		{"not found by nil tree", []int{}, 1, []int{}, false},
		{"not found only root", []int{1}, 2, []int{1}, false},
		{
			"not found",
			[]int{5, 7, 2, 1, 4, 6},
			3,
			[]int{1, 2, 4, 5, 6, 7},
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)

			result := bst.Remove(tc.deleteValue)
			if result != tc.resultExpected {
				t.Errorf("Expected delete result to be '%v'. Got '%v'", tc.resultExpected, result)
			}

			var inOrderResult []int
			bst.InOrderTraverse(func(i int) {
				inOrderResult = append(inOrderResult, i)
			})

			if bst.Len() != len(tc.inOrderExpected) {
				t.Errorf("Expected Len value to be '%v'. Got '%v'", len(tc.inOrderExpected), bst.Len())
			}

			if len(inOrderResult) != len(tc.inOrderExpected) {
				t.Fatalf("Expected in order traversal to be %v. Got %v", tc.inOrderExpected, inOrderResult)
			}

			for i := range inOrderResult {
				if inOrderResult[i] != tc.inOrderExpected[i] {
					t.Errorf("Expected in order traversal after remove to be '%v'. Got '%v'", tc.inOrderExpected, inOrderResult)
					break
				}
			}
		})
	}
}

func TestPrint(t *testing.T) {
	tt := []struct {
		name         string
		insertValues []int
		result       string
	}{
		{"empty tree", []int{}, ""},
		{"only root tree", []int{1}, "-[1\n"},
		{
			"plain tree",
			[]int{2, 1, 3},
			fmt.Sprintf("%s-[3\n-[2\n%[1]s-[1\n", trees.PrintLevelSeparator),
		},
		{
			"unbalanced to right",
			[]int{5, 6, 7, 8, 9},
			fmt.Sprintf(
				"%s%[1]s%[1]s%[1]s-[9\n%[1]s%[1]s%[1]s-[8\n%[1]s%[1]s-[7\n%[1]s-[6\n-[5\n",
				trees.PrintLevelSeparator),
		},
		{
			"unbalanced to left",
			[]int{5, 4, 3, 2, 1},
			fmt.Sprintf(
				"-[5\n%s-[4\n%[1]s%[1]s-[3\n%[1]s%[1]s%[1]s-[2\n%[1]s%[1]s%[1]s%[1]s-[1\n",
				trees.PrintLevelSeparator),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bst := trees.BST{}
			fillTreeWithList(&bst, tc.insertValues)
			buf := new(bytes.Buffer)
			bst.Print(buf)

			if buf.String() != tc.result {
				t.Errorf("Expected print to be:\n%v\nGot:\n%v", tc.result, buf.String())
			}
		})
	}

}
