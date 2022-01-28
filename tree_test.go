package RBTree

import (
	"strconv"
	"testing"
)

func TestCreateTree(t *testing.T) {
	root := &Node{value: NodeVal(2)}
	root.Insert(&Node{value: NodeVal(1)})
	root.Insert(&Node{value: NodeVal(3)})
	root.Insert(&Node{value: NodeVal(0)})
	PrintTree(root)
}

type NodeVal int

func (nodeVal NodeVal) compare(in Value) bool {
	return int(in.(NodeVal)) > int(nodeVal)
}

func (nodeVal NodeVal) String() string {
	return strconv.Itoa(int(nodeVal))
}
