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
		if len(val) != len(node.value) {
			return false
		}
		matchLength := 0
		for i := 0; i < len(val); i++ {
			if val[i] != node.value[i] {
				break
			}
			matchLength++
		}
		if matchLength == len(val) {
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
		for i := 0; i < commonLength; i++ {
			if val[i] != node.value[i] {
				goto A
			}
		}
		if commonLength > maxMatch {
			maxMatch = commonLength
		}

	A:
		{
		}

		node = node.next
	}

	return maxMatch
}
