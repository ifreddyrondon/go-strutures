package binarytrees

// BNode is a single node that compose a binary tree.
type BNode struct {
	Value int
	Left  *BNode
	Right *BNode
}

// NewBNode is a helper function that given a value return a node.
func NewBNode(value int) *BNode {
	return &BNode{value, nil, nil}
}
