package RBTree

import "errors"

type Node struct {
	value   Value
	left    *Node
	right   *Node
	isBlack bool
}

type Value interface {
	compare(in Value) bool //in > val -> true
	String() string
}

func (node *Node) insertion(ins *Node) {
	if node.value.compare(ins.value) {
		if node.right == nil {
			node.right = ins
		} else {
			node.right.insertion(ins)
		}
	} else {
		if node.left == nil {
			node.left = ins
		} else {
			node.left.insertion(ins)
		}
	}
}

func (node *Node) Insert(ins *Node) error {
	err := checkInserting(ins)
	if err != nil {
		return err
	}
	node.insertion(ins)

	return nil
}

func checkInserting(ins *Node) error {
	if ins.isBlack || ins.left != nil || ins.right != nil {
		return errors.New("inserting node must be red and have all children nil")
	}
	return nil
}
