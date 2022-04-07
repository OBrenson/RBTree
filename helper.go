package RBTree

import (
	"container/list"
	"fmt"
	"math"
)

var (
	red   = "\033[31m"
	reset = "\033[0m"
	arrow = `|`
)

/*
Print tree in a terminal. Very helpful for testing
*/
func (t *RBTree) PrintTree() {
	h := getHeight(t.root, 0)
	copyNode := &Node{Key: t.root.Key, isBlack: t.root.isBlack}
	copyTree(t.root, copyNode, 1, h, false)
	tabs := calcTabs(h)
	printTree(copyNode, tabs)
}

func printTree(root *Node, tabs []string) {
	l := list.New()
	l.PushBack(root)
	level := 0
	indLevel := 0
	for l.Len() != 0 {
		el := l.Front()
		n := el.Value.(*Node)
		l.Remove(el)
		if n.Key != nil {
			if n.isBlack {
				fmt.Print(tabs[indLevel] + n.Key.String())
			} else {
				fmt.Print(tabs[indLevel] + red + n.Key.String() + reset)
			}
		} else {
			fmt.Print(tabs[indLevel] + "  ")
		}
		if n.left != nil {
			l.PushBack(n.left)
		}
		if n.right != nil {
			l.PushBack(n.right)
		}
		level++
		if level == int(math.Pow(2, float64(indLevel))) {
			fmt.Println()
			level = 0
			indLevel++
		}
	}
}

func calcTabs(height int) (res []string) {
	tab := " "
	tmpTab := ""
	for i := height; i > 0; i-- {
		for j := 0; j < int(math.Pow(2, float64(i)))/2+2; j++ {
			tmpTab += tab
		}
		res = append(res, tmpTab)
		tmpTab = ""
	}
	return res
}

func getHeight(node *Node, height int) int {
	if node == nil {
		return height
	}
	heightLeft := getHeight(node.left, height+1)
	heightRight := getHeight(node.right, height+1)
	if heightLeft < heightRight {
		return heightRight
	}
	return heightLeft
}

func copyTree(inN, prevOutN *Node, curH, h int, isLeft bool) {
	if curH == h+1 {
		return
	}
	if curH != 1 {
		var nNode *Node
		if inN == nil {
			nNode = &Node{Key: nil}
		} else {
			nNode = &Node{Key: inN.Key, isBlack: inN.isBlack}
		}
		if isLeft {
			prevOutN.left = nNode
		} else {
			prevOutN.right = nNode
		}
		if inN != nil {
			copyTree(inN.left, nNode, curH+1, h, true)
		} else {
			copyTree(nil, nNode, curH+1, h, true)
		}
		if inN != nil {
			copyTree(inN.right, nNode, curH+1, h, false)
		} else {
			copyTree(nil, nNode, curH+1, h, false)
		}
	} else {
		copyTree(inN.left, prevOutN, curH+1, h, true)
		copyTree(inN.right, prevOutN, curH+1, h, false)
	}
}
