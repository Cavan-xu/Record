package forbid

type List struct {
	tail *Node
}

type Node struct {
	value []rune
	prev  *Node
	next  *Node
}

func NewList() *List {
	node := &Node{}
	node.prev = node
	node.next = node
	return &List{tail: node}
}

func (list *List) Add(node *Node) {
	node.next = list.tail
	node.prev = list.tail.prev
	list.tail.prev.next = node
	list.tail.prev = node
}

func (list *List) GetFirst() *Node {
	return list.tail.next
}

func (list *List) ExactMatchSearch(val []rune) bool {
	node := list.GetFirst()

	for node != list.tail {
		if string(node.value) == string(val) {
			return true
		}
		node = node.next
	}

	return false
}

func (list *List) CommonPrefixSearch(val []rune) int {
	maxMatch := -1
	node := list.GetFirst()

	for node != list.tail {
		commonLength := len(node.value)
		if len(val) < len(node.value) {
			node = node.next
			continue
		}
		if string(val[:commonLength]) == string(node.value) && commonLength > maxMatch {
			maxMatch = commonLength
		}
		node = node.next
	}

	return maxMatch
}
