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

// Insert insert an item in the right position in the tree. Return true if the value was inserted and false otherwise
func (t *BST) Insert(value int) bool {
	node := NewNode(value)
	if t.Root == nil {
		t.Root = node
		return true
	}
	return insertNode(t.Root, node)
}

func insertNode(root, newNode *Node) bool {
	if root.Value == newNode.Value {
		return false
	}

	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
			return true
		}
		return insertNode(root.Left, newNode)
	} else {
		if root.Right == nil {
			root.Right = newNode
			return true
		}
		return insertNode(root.Right, newNode)
	}
}

// InOrderTraverse visits all the nodes in order
func (t *BST) InOrderTraverse(f func(int)) {
	inOrderTraverse(t.Root, f)
}

func inOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}

	inOrderTraverse(node.Left, f)
	f(node.Value)
	inOrderTraverse(node.Right, f)
}

// PreOrderTraverse visits all the nodes in pre order
func (t *BST) PreOrderTraverse(f func(int)) {
	preOrderTraverse(t.Root, f)
}

func preOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}

	f(node.Value)
	preOrderTraverse(node.Left, f)
	preOrderTraverse(node.Right, f)
}

// PostOrderTraverse visits all the nodes in post order
func (t *BST) PostOrderTraverse(f func(int)) {
	postOrderTraverse(t.Root, f)
}

func postOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}

	postOrderTraverse(node.Left, f)
	postOrderTraverse(node.Right, f)
	f(node.Value)
}

// Min returns the node with minimal value stored in the tree
func (t *BST) Min() *Node {
	return minNode(t.Root)
}

func minNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	current := node
	for {
		if current.Left == nil {
			return current
		}
		current = current.Left
	}
}

// Max returns the node with maximum value stored in the tree
func (t *BST) Max() *Node {
	return maxNode(t.Root)
}

func maxNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	current := node
	for {
		if current.Right == nil {
			return current
		}
		current = current.Right
	}
}

// Search returns the node if the value exists in the tree
func (t *BST) Search(value int) *Node {
	return searchNode(t.Root, value)
}

func searchNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if node.Value > value {
		return searchNode(node.Left, value)
	} else {
		return searchNode(node.Right, value)
	}
}

// Has returns true if if the value exists in the tree
func (t *BST) Has(value int) bool {
	return t.Search(value) != nil
}

// Remove remove an item from the tree. Return true if the value was removed and false otherwise.
func (t *BST) Remove(value int) bool {
	_, removed := removeNode(t.Root, value)
	return removed
}

func removeNode(node *Node, value int) (*Node, bool) {
	removed := true
	if node == nil {
		return nil, false
	}

	if node.Value > value {
		node.Left, removed = removeNode(node.Left, value)
		return node, removed
	} else if node.Value < value {
		node.Right, removed = removeNode(node.Right, value)
		return node, removed
	}

	// after this point the node.Value == value
	// delete case 1: delete leaf node
	if node.Left == nil && node.Right == nil {
		return nil, removed
	}

	// delete case 2: delete half-leaf node
	if node.Left == nil {
		node = node.Right
		return node, removed
	}
	if node.Right == nil {
		node = node.Left
		return node, removed
	}

	// delete case 3: delete an inner node
	replacement := maxNode(node.Left)
	node.Value = replacement.Value
	node.Left, removed = removeNode(node.Left, node.Value)
	return node, removed
}
