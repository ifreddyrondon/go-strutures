package trees

import "math/rand"

// NewRandBST returns a new, random binary search tree
func NewRandBST(length int) *BST {
	t := &BST{}
	for _, v := range rand.Perm(length) {
		t.Insert(1 + v)
	}
	return t
}
