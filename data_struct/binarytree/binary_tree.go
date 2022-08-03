package binarytree

import (
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

func (t *Tree) PreOrder(node *Node, res []int) []int {
	if node == nil {
		return res
	}

	res = append(res, node.val)
	res = t.PreOrder(node.left, res)
	res = t.PreOrder(node.right, res)
	return res
}

// 使用递归打印中序遍历
func (t *Tree) InOrder(node *Node, res []int) []int {
	if node == nil {
		return res
	}

	res = t.InOrder(node.left, res)
	res = append(res, node.val)
	res = t.InOrder(node.right, res)
	return res
}

func (t *Tree) AfterOrder(node *Node, res []int) []int {
	if node == nil {
		return res
	}

	res = t.AfterOrder(node.left, res)
	res = t.AfterOrder(node.right, res)
	res = append(res, node.val)
	return res
}

func (t *Tree) PreOrderV2() []int {
	res := make([]int, 0, t.count)
	stack := make([]*Node, 0, t.count)
	stack = append(stack, t.root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.val)
		if node.right != nil {
			stack = append(stack, node.right)
		}
		if node.left != nil {
			stack = append(stack, node.left)
		}
	}

	return res
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

func (t *Tree) AfterOrderV2() []int {
	var prev *Node
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
		if node.right == nil || node.right == prev {
			res = append(res, node.val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.right
		}
	}

	return res
}
