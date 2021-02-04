package test

import (
	"algorithm/BST"
	"fmt"
	"testing"
)

func TestBSt(t *testing.T) {
	bst := BST.NewBST()
	nums := []int{5, 3, 6, 8, 4, 2}
	for i := 0; i < len(nums); i++ {
		bst.Add(nums[i])
	}
	bst.LevelOrder()
	fmt.Println(bst)
}
