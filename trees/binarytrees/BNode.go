package binarytrees

// BNode is a single node that compose a binary tree.
type BNode struct {
	Value int
	Left  *BNode
	Right *BNode
}

// NewNode is a helper function that given a value return a node.
func NewNode(value int) *BNode {
	return &BNode{value, nil, nil}
}
