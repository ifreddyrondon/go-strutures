package binarytrees

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
)

const (
	PrintLevelSeparator = "       "
	PrintNode           = "-["
)

// NewRandBST returns a new, random binary search tree
func NewRandBST(length int) *BST {
	t := &BST{}
	for _, v := range rand.Perm(length) {
		t.Insert(1 + v)
	}
	return t
}

// PrintTreeFromNode prints a visual representation of the binary tree from a given node into an io.Writer
func PrintTreeFromNode(w io.Writer, n *BNode, level int) {
	if n != nil {
		format := bytes.NewBufferString("")
		for i := 0; i < level; i++ {
			format.WriteString(PrintLevelSeparator)
		}
		format.WriteString(PrintNode)
		level++
		PrintTreeFromNode(w, n.Right, level)
		fmt.Fprintf(w, "%s%d\n", format.String(), n.Value)
		PrintTreeFromNode(w, n.Left, level)
	}
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
