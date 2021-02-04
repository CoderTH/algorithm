package BST

import (
	"algorithm/Queue"
	"fmt"
)

type node struct {
	E     int
	Left  *node
	Right *node
}
type BST struct {
	root *node
	size int
}

func NewNode(E int) *node {
	return &node{
		E:     E,
		Left:  nil,
		Right: nil,
	}
}
func NewBST() *BST {
	return &BST{
		root: nil,
		size: 0,
	}
}
func (bst *BST) Size() int {
	return bst.size
}
func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

//递归实现
func (bst *BST) Add(e int) {
	bst.root = bst.add(bst.root, e)
}
func (bst *BST) add(Node *node, e int) *node {
	if Node == nil {
		bst.size++
		return NewNode(e)
	}
	if Node.E > e {
		Node.Left = bst.add(Node.Left, e)
	} else if Node.E < e {
		Node.Right = bst.add(Node.Right, e)
	}

	return Node
}

//非递归实现
func (bst *BST) Add2(e int) {
	if bst.root == nil {
		bst.root = NewNode(e)
		bst.size++
		return
	}
	p := bst.root
	for p != nil {
		//当前节点大于带插入节点
		//如果其左子树为空那么将插入其左子树，否则将p下移至其左子树
		if p.E > e {
			if p.Left == nil {
				p.Left = NewNode(e)
				bst.size++
				return
			}
			p = p.Left
			//当前节点小于代插入节点
			//如果其右子树为空那么将插入其右子树，否则将p下移至其右子树
		} else if p.E < e {
			if p.Right == nil {
				p.Right = NewNode(e)
				bst.size++
				return
			}
			p = p.Right
			//这里实现的二分搜索树中是不存在相同元素的，所以当遇到了相同的元素就直接返回
		} else {
			return
		}
	}
}

//递归实现（查找）
func (bst *BST) Contains(e int) bool {
	return bst.contains(bst.root, e)
}
func (bst *BST) contains(Node *node, e int) bool {
	if Node == nil {
		return false
	}
	if Node.E == e {
		return true
	} else if Node.E > e {
		return bst.contains(Node.Left, e)
	} else {
		return bst.contains(Node.Right, e)
	}
}

//前序遍历
func (bst *BST) PreOrder() {
	bst.preOrder(bst.root)
}
func (bst *BST) preOrder(Node *node) {
	if Node == nil {
		return
	}
	fmt.Println(Node.E)
	bst.preOrder(Node.Left)
	bst.preOrder(Node.Right)
}

//中序遍历
func (bst *BST) InOrder() {
	bst.inOrder(bst.root)
}
func (bst *BST) inOrder(Node *node) {
	if Node == nil {
		return
	}
	bst.inOrder(Node.Left)
	fmt.Println(Node.E)
	bst.inOrder(Node.Right)
}

//后序遍历
func (bst *BST) PostOrder() {
	bst.postOrder(bst.root)
}
func (bst *BST) postOrder(Node *node) {
	if Node == nil {
		return
	}
	bst.postOrder(Node.Left)
	bst.postOrder(Node.Right)
	fmt.Println(Node.E)
}

//层序遍历
func (bst *BST) LevelOrder() {
	queue := Queue.NewLinkedListQueue()
	queue.Enqueue(bst.root)
	for !queue.IsEmpty() {
		cur := queue.Dequeue().(*node)
		fmt.Println(cur.E)
		if cur.Left != nil {
			queue.Enqueue(cur.Left)
		}
		if cur.Right != nil {
			queue.Enqueue(cur.Right)
		}
	}
}
//寻找二分搜索树中最小的元素
func (bst *BST) MiniMun()int {
	if bst.size == 0 {
		panic("BST is empty!")
	}
	return bst.minMun(bst.root).E
}

func (bst *BST) minMun(Node *node)*node {
	if Node.Left == nil {
		return Node
	}
	return bst.minMun(Node.Left)

}
//寻找二分搜索树中最大的元素
func (bst *BST) MaxiMun()int {
	if bst.size == 0 {
		panic("BST is empty!")
	}
	return bst.maxMun(bst.root).E
}

func (bst *BST) maxMun(Node *node)*node {
	if Node.Right == nil {
		return Node
	}
	return bst.maxMun(Node.Right)

}
//从二分搜索树中删除最小值虽在的节点，返回最小值
func (bst *BST) RemoveMin() int{
	ret := bst.MiniMun()
	bst.root= bst.removeMin(bst.root)
	return ret

}
//删除以node为根的二分搜索树中的最小节点
//返回删除节点后新的二分搜索树的根
func (bst *BST) removeMin(Node *node) *node{
	if Node.Left == nil{
		rightNode := Node.Right
		Node.Right = nil
		bst.size--
		return rightNode
	}
	Node.Left= bst.removeMin(Node.Left)
	return Node
}



//从二分搜索树中删除最小值虽在的节点，返回最小值
func (bst *BST) RemoveMax() int{
	ret := bst.MaxiMun()
	bst.root= bst.removeMax(bst.root)
	return ret

}
//删除以node为根的二分搜索树中的最小节点
//返回删除节点后新的二分搜索树的根
func (bst *BST) removeMax(Node *node) *node{
	if Node.Right == nil{
		leftNode := Node.Left
		Node.Left = nil
		bst.size--
		return leftNode
	}
	Node.Right= bst.removeMin(Node.Right)
	return Node
}

func (bst *BST) Remove(E int) {
	bst.root = bst.remove(bst.root,E)
}
func (bst *BST) remove(Node *node,E int) *node {
	if Node == nil {
		return nil
	}
	if E < Node.E {
		Node.Left= bst.remove(Node.Left, E)
		return Node
	}else  if E > Node.E {
		Node.Right= bst.remove(Node.Right, E)
		return Node
	}else {
		//待删除节点左子树为空的情况
		if Node.Left == nil{
			rightNode := Node.Right
			Node.Right = nil
			bst.size--
			return rightNode
		}
		//待删除节点右子树为空的情况
		if Node.Right == nil{
			leftNode := Node.Left
			Node.Left = nil
			bst.size--
			return leftNode
		}
		//待删除节点左右子树都为空的情况
		//待着比待删除节点大的最小节点，即待删除节点右子树最小的节点
		//用这个节点顶替待删除节点的位置
		successor := bst.minMun(Node.Right)
		successor.Right = bst.removeMin(Node.Right)
		bst.size++
		successor.Left = Node.Left
		Node=nil
		bst.size--
		return successor
	}
}

//func(bst *BST)String()string  {
//	var res string
//	res = bst.generateBSTString(bst.root,0,res)
//	return res
//}
//
//func (bst *BST)generateBSTString(Node *node,depth int,res string) string {
//	if Node == nil {
//		fmt.Sprintf("%S%Snil\n",res,bst.generateDepthString(depth))
//		return res
//	}
//	fmt.Sprintf("%s%s%d\n",res,bst.generateDepthString(depth),Node.E)
//	res = bst.generateBSTString(Node.Left,depth+1,res)
//	res = bst.generateBSTString(Node.Right,depth+1,res)
//	return res
//}
//func (bst *BST)generateDepthString(depth int) string {
//	var str string
//	for i:=1;i<depth;i++{
//		str = fmt.Sprintf("%s--",str)
//	}
//	return str
//}
//
