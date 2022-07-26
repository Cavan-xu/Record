package balancetree

/*
	平衡二叉树，又称 AVL 树，可以是一棵空树，或者是具有以下性质的二叉排序树：
		1.左子树和右子树的高度之差的绝对值不超过1
		2.它的左子树和右子树都是一颗平衡二叉树

	引入原因：为了解决排序二叉树退化成链表的情况
*/

type Node struct {
	Val        int
	height     int   // 节点高度
	depth      int   // 节点深度
	leftChild  *Node // 左节点
	rightChild *Node // 右节点
}
