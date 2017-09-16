package trees

// Node is a single node that compose a tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewNode is a helper function that given a value return a node.
func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}

type Tree interface {
	Root() *Node
}
