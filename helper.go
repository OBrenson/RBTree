package RBTree

import (
	"container/list"
	"fmt"
)

var (
	red   = "\033[31m"
	reset = "\033[0m"
	arrow = `|`
)

func PrintTree(root *Node) {

	res := printTree(root)

	for _, nodes := range res {
		for _, node := range nodes {
			fmt.Print(node.value.String())
		}
		fmt.Println()
	}
}

func printTree(root *Node) [][]*Node {
	nodesOnLevels := make([][]*Node, 0)
	nodes := make([]*Node, 0)

	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		el := l.Front()
		l.Remove(el)

		if len(nodes) != 0 && (el.Value.(*Node) == nodes[0].left || el.Value.(*Node) == nodes[0].right) {
			nodesOnLevels = append(nodesOnLevels, nodes)
			nodes = make([]*Node, 0)
		}

		if el.Value.(*Node).left != nil {
			l.PushBack(el.Value.(*Node).left)
		}
		if el.Value.(*Node).right != nil {
			l.PushBack(el.Value.(*Node).right)
		}

		nodes = append(nodes, el.Value.(*Node))
	}
	if len(nodes) != 0 {
		nodesOnLevels = append(nodesOnLevels, nodes)
	}
	return nodesOnLevels
}
