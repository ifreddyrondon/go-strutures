package trees_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ifreddyrondon/go-strutures/trees"
)

func TestNewRandBST(t *testing.T) {
	bst := trees.NewRandBST(10)

	if bst.Root == nil {
		t.Error("Expected root to be not nil")
	}

	if bst.Len() != 10 {
		t.Errorf("Expected Len value to be '%v'. Got '%v'", 10, bst.Len())
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
			trees.Print(buf, bst)

			if buf.String() != tc.result {
				t.Errorf("Expected print to be:\n%v\nGot:\n%v", tc.result, buf.String())
			}
		})
	}

}
