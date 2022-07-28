package binarytree

import (
	"fmt"

	"awesomeProject/common"
)

/*
	二叉树：是指树中节点的度不大于2的有序树，有以下规定：
		1.可以是一棵空树，或者是一颗由根节点和两颗互不相交的，被称为左子树和右子树组成的非空树
		2.左子树和右子树同样都是二叉树

	二叉树的构建：确定一颗二叉树，需要通过二叉树的前序遍历、中序遍历，或者是后序遍历、中序遍历来确定

	如下二叉树：
				8
		5				9
	1		6					12
							11		13

	中序遍历: 左根右: {1,5,6,8,9,11,12,13}	中序遍历得到的是排好序的
	先序遍历: 根左右: {8,5,1,6,9,12,11,13}
	后序遍历: 左右根: {1,6,5,11,13,12,9,8}
*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	root  *Node
	count int
}

func NewTree(inOrder, preOrder []int) *Tree {
	root := generateNode(inOrder, preOrder)

	return &Tree{root: root, count: len(inOrder)}
}

// 使用递归构造子树
func generateNode(inOrder, preOrder []int) *Node {
	if len(inOrder) == 0 {
		return nil
	}

	node := &Node{val: preOrder[0]}
	index := common.IntIndexOf(preOrder[0], inOrder)

	left := generateNode(inOrder[:index], preOrder[1:index+1])
	right := generateNode(inOrder[index+1:], preOrder[index+1:])

	node.left = left
	node.right = right

	return node
}

func (t *Tree) PreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.val)
	t.PreOrder(node.left)
	t.PreOrder(node.right)
}

// 使用递归打印中序遍历
func (t *Tree) InOrder(node *Node) {
	if node == nil {
		return
	}

	t.InOrder(node.left)
	fmt.Println(node.val)
	t.InOrder(node.right)
}

func (t *Tree) AfterOrder(node *Node) {
	if node == nil {
		return
	}

	t.AfterOrder(node.left)
	t.AfterOrder(node.right)
	fmt.Println(node.val)
}

func (t *Tree) InOrderV2() []int {
	node := t.root
	res := make([]int, 0, t.count)
	stack := make([]*Node, 0, t.count)

	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.val)
		node = node.right
	}

	return res
}
