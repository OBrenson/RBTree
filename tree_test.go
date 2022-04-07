package RBTree

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

type NodeVal int

func (nodeVal NodeVal) Compare(in Key) bool {
	return int(in.(NodeVal)) > int(nodeVal)
}

func (nodeVal NodeVal) String() string {
	return strconv.Itoa(int(nodeVal))
}

func (nodeVal NodeVal) Equal(in Key) bool {
	return int(in.(NodeVal)) == int(nodeVal)
}

/*
Simple test without any automation. Just create tree, insert some values
and see the result in the terminal with your own eyes
*/
func TestCreateTree(t *testing.T) {

	tree := createTestTree()

	fmt.Println()
	tree.PrintTree()
}

func TestInsert(t *testing.T) {
	tree := CreateTree(NodeVal(0))
	for i := 0; i < 40; i++ {
		counter := 0
		err := tree.Insert(NodeVal(rand.Int31n(100)), nil)
		if err != nil {
			t.Error(err)
			tree.PrintTree()
			return
		}
		if err = checkDepths(tree.root, &counter, 0, true); err != nil {
			t.Error(err)
			tree.PrintTree()
			return
		}
		if err = checkParents(tree.root); err != nil {
			t.Error(err)
			tree.PrintTree()
			return
		}
	}
	tree.PrintTree()
	counter := 0
	if err := checkDepths(tree.root, &counter, 0, true); err != nil {
		t.Error(err)
	}

	if err := tree.InsertNode(&Node{Key: nil}); err != nil {
		t.Log(err)
	} else {
		t.Error("must check is key has nil value")
	}

	v := "testValue"
	_ = tree.Insert(NodeVal(0), v)
	n := tree.Find(NodeVal(0))
	if n.Value != v {
		t.Error("node must has Value " + v)
	}
}

func TestFind(t *testing.T) {
	tree := createTestTree()

	node := tree.Find(NodeVal(9))
	if node == nil || int(node.Key.(NodeVal)) != 9 {
		t.Error("could not find node")
	}
	node = tree.Find(NodeVal(-100))
	if node != nil {
		t.Error("this node should not be")
	}
	node = tree.Find(NodeVal(100))
	if node != nil {
		t.Error("this node should not be")
	}
}

func TestRemove(t *testing.T) {
	nodes := [100]NodeVal{13, 90, 94, 63, 33, 47, 78, 24, 59, 53, 57, 21, 89, 99, 0, 5, 88, 38, 3, 55, 51, 10, 5, 56, 66, 28, 61, 2, 83, 46, 63, 76, 2, 18, 47, 94, 77, 63, 96, 20, 23, 53, 37, 33, 41, 59, 33, 43, 91, 2, 78, 36, 46, 7, 40, 3, 52, 43, 5, 98, 25, 51, 15, 57, 87, 10, 10, 85, 90, 32, 98, 53, 91, 82, 84, 97, 67, 37, 71, 94, 26, 2, 81, 79, 66, 70, 93, 86, 19, 81, 52, 75, 85, 10, 87, 49, 28, 18, 84, 3}

	tree := CreateTree(NodeVal(0))
	capac := len(nodes)

	for i := 0; i < capac; i++ {
		err := tree.Insert(nodes[i], nil)
		if err != nil {
			t.Error(err)
			tree.PrintTree()
			return
		}
	}
	for i := 0; i < capac; i++ {
		fmt.Println(i)
		nfErr := tree.Remove(nodes[i])
		if nfErr != nil {
			t.Log(nfErr)
		}
		counter := 0
		err := checkParents(tree.root)
		if err != nil {
			t.Error(err)
			tree.PrintTree()
			return
		}
		err = checkDepths(tree.root, &counter, 0, true)
		if err != nil {
			t.Error(err)
			tree.PrintTree()
			return
		}
		if nfErr == nil {
			t.Log("all ok\n")
		}
	}
}

func TestGetAll(t *testing.T) {
	tree := CreateTree(NodeVal(0))
	arr := make([]int, 10)
	for i := 1; i < 11; i++ {
		_ = tree.Insert(NodeVal(i), nil)
		arr = append(arr, i)
	}
LOOP:
	for _, node := range tree.GetAll() {
		var v int
		for v = range arr {
			if v == int(node.Key.(NodeVal)) {
				continue LOOP
			}
		}
		t.Error("GetAll must return value " + strconv.Itoa(v))
	}
}

func TestGetAllSorted(t *testing.T) {
	tree := Create()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		_ = tree.Insert(NodeVal(arr[i]), nil)
	}
	res := tree.GetSorted()
	for i := 0; i < len(res); i++ {
		if int(res[i].Key.(NodeVal)) != arr[i] {
			t.Error("arrays don`t equal")
		}
	}
}

func insertValue(tree *RBTree, key Key) {
	err := tree.InsertNode(&Node{Key: key})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func createTestTree() *RBTree {
	tree := Create()
	insertValue(tree, NodeVal(3))
	insertValue(tree, NodeVal(4))
	insertValue(tree, NodeVal(2))
	insertValue(tree, NodeVal(0))
	insertValue(tree, NodeVal(7))
	insertValue(tree, NodeVal(8))
	insertValue(tree, NodeVal(9))
	insertValue(tree, NodeVal(10))
	insertValue(tree, NodeVal(11))
	insertValue(tree, NodeVal(-2))
	insertValue(tree, NodeVal(-5))
	insertValue(tree, NodeVal(-8))
	return tree
}

//checking nodes` pointers violations
func checkParents(node *Node) error {
	if node.parent != nil {
		if node != node.parent.left && node != node.parent.right {
			return errors.New("PARENT ERROR")
		}
	}
	var err error
	if node.left.Key != nil {
		err = checkParents(node.left)
	}
	if err != nil {
		return err
	}
	if node.right.Key != nil {
		err = checkParents(node.right)
	}
	return err
}

//checking depth of black nodes in each subtree
func checkDepths(n *Node, overall *int, cur int, pIsB bool) error {
	if n.isBlack {
		cur++
	}
	if n.left.Key == nil || n.right.Key == nil {
		if *overall == 0 {
			*overall = cur
		} else {
			if *overall != cur {
				return errors.New("black depth of all subtrees must be equal")
			}
		}
		if !n.isBlack && !pIsB {
			return errors.New("parent and child are red " + n.Key.String())
		}
	}
	var err error = nil
	if n.left.Key != nil {
		err = checkDepths(n.left, overall, cur, n.isBlack)
	} else {
		if !n.left.isBlack {
			return errors.New("wrong color of nil node")
		}
	}
	if err != nil {
		return err
	}
	if n.right.Key != nil {
		err = checkDepths(n.right, overall, cur, n.isBlack)
	} else {
		if !n.right.isBlack {
			return errors.New("wrong color of nil node")
		}
	}
	return err
}
