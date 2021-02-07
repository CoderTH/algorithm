package Set

import "algorithm/BST"

type BSTSet struct {
	bst *BST.BST
}

func NewSet() BSTSet {
	return BSTSet{
		bst: BST.NewBST(),
	}
}

func (B *BSTSet) Add(e int) {
	B.bst.Add(e)
}

func (B *BSTSet) Remove(e int) {
	B.bst.Remove(e)
}

func (B *BSTSet) Contains(e int) bool {
	B.bst.Contains(e)
}

func (B *BSTSet) GetSize() int {
	return B.bst.Size()
}

func (B *BSTSet) IsEmpty() bool {
	return B.bst.IsEmpty()
}
