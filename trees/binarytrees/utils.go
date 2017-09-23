package binarytrees

import (
	"fmt"
	"io"
	"math/rand"

	"bytes"

	"github.com/ifreddyrondon/gostrutures"
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

// PrintTreeByLevel prints a visual representation by levels of a tree from a given node into an io.Writer
func PrintTreeByLevel(w io.Writer, n *BNode) {
	if n == nil {
		return
	}

	queue := gostrutures.Queue{}
	queue.Push(n)
	nodesInCurrentLevel, nodesInNextLevel := 1, 0

	for {
		if queue.Size() == 0 {
			break
		}
		node := queue.Pop().(*BNode)
		fmt.Fprintf(w, "%v ", node.Value)
		nodesInCurrentLevel--
		if node.Left != nil {
			queue.Push(node.Left)
			nodesInNextLevel += 1
		}
		if node.Right != nil {
			queue.Push(node.Right)
			nodesInNextLevel += 1
		}
		if nodesInCurrentLevel == 0 {
			fmt.Fprintln(w)
			nodesInCurrentLevel = nodesInNextLevel
			nodesInNextLevel = 0
		}
	}
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
