package trees

// A single node that compose a tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}

// Binary search tree
type BST struct {
	Root *Node
}

// Constructor, return a Binary Search Tree with the root
func New(value int) *BST {
	return &BST{NewNode(value)}
}

// Insert insert an item in the right position in the tree
func (t *BST) Insert(value int) {
	node := NewNode(value)
	if t.Root == nil {
		t.Root = node
	} else {
		insertNode(t.Root, node)
	}
}

func insertNode(root, newNode *Node) {
	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
		} else {
			insertNode(root.Left, newNode)
		}
	} else {
		if root.Right == nil {
			root.Right = newNode
		} else {
			insertNode(root.Right, newNode)
		}
	}
}
