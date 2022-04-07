package RBTree

import (
	"errors"
)

/*
	RBTree - Black red tree realisation. All last nodes have black leaves with the nil Key.
			Algorithm of insertion has been taken from here https://habr.com/ru/company/otus/blog/472040/
			Algorithm of remove has been taken from here https://en.wikipedia.org/wiki/Red–black_tree
*/
type RBTree struct {
	root *Node
}

/*
	Node - standard node of rb tree. All logic is tied up in the keys comparison,
		because of this key must realise Key interface and should not be nil.
		In Value you can save whatever you want.
*/
type Node struct {
	Key     Key
	left    *Node
	right   *Node
	parent  *Node
	isBlack bool
	Value   interface{}
}

type Key interface {
	Compare(in Key) bool // in > val -> true
	Equal(in Key) bool   // in == val -> true
	String() string
}

/*
	Create root of rb tree.
*/
func CreateTree(key Key) *RBTree {
	root := &Node{Key: key, isBlack: true}
	root.left = &Node{Key: nil, isBlack: true, parent: root}
	root.right = &Node{Key: nil, isBlack: true, parent: root}
	return &RBTree{root: root}
}

/*
	Create tree without root
*/
func Create() *RBTree {
	return &RBTree{}
}

/*
	Insert new Node in tree. If Node all ready exists Node.Value will be changed
*/
func (t *RBTree) InsertNode(ins *Node) error {

	err := checkInserting(ins)
	if err != nil {
		return err
	}
	ins.left = &Node{Key: nil, isBlack: true, parent: ins}
	ins.right = &Node{Key: nil, isBlack: true, parent: ins}
	if t.root == nil {
		t.root = ins
		return nil
	}
	if insertion(ins, t.root) {
		balancing(ins)
		t.root = getRoot(ins)
		t.root.isBlack = true
	}
	return nil
}

/*
	Wrapper for InsertNode. Create new Node from key and value
*/
func (t *RBTree) Insert(key Key, value interface{}) error {
	ins := &Node{Key: key, Value: value}

	return t.InsertNode(ins)
}

/*
	Find node by its Key
*/
func (t *RBTree) Find(key Key) *Node {
	return find(key, t.root)
}

/*
	Remove node by its Key, if the node is single in the tree, it can`t be removed
*/
func (t *RBTree) Remove(key Key) error {
	node := find(key, t.root)
	if node == nil {
		return errors.New("no such element")
	}
	if node == t.root {
		return nil
	}
	maxL := findMaxLeft(node)

	maxL.Key, node.Key = node.Key, maxL.Key
	bal := deleting(maxL)
	delBalancing(bal)
	t.root = getRoot(node)
	return nil
}

/*
	Return all nodes of the tree. Does not guarantee the sorted sequence
*/
func (t *RBTree) GetAll() []*Node {
	nodes := make([]*Node, 0)
	createNodesSlice(&nodes, t.root)
	return nodes
}

/*
	Return all nodes of the tree. Guarantee the sorted sequence
*/
func (t *RBTree) GetSorted() []*Node {
	res := make([]*Node, 0)
	maxLeft := t.root
	for maxLeft.left.Key != nil {
		maxLeft = maxLeft.left
	}
	createSortedSlice(&res, maxLeft)
	return res
}

func createSortedSlice(nodes *[]*Node, node *Node) {
	*nodes = append(*nodes, node)
	if node.right.Key != nil {
		maxLeft := node.right
		for maxLeft.left.Key != nil {
			maxLeft = maxLeft.left
		}
		createSortedSlice(nodes, maxLeft)
	}
	if node.parent != nil && node.parent.left == node {
		createSortedSlice(nodes, node.parent)
	}
}

func createNodesSlice(nodes *[]*Node, node *Node) {
	if node.Key == nil {
		return
	}
	*nodes = append(*nodes, node)
	createNodesSlice(nodes, node.left)
	createNodesSlice(nodes, node.right)
}

func find(key Key, node *Node) *Node {
	if node.Key.Equal(key) {
		return node
	}
	if node.Key.Compare(key) {
		if node.right.Key == nil {
			return nil
		}
		return find(key, node.right)
	} else {
		if node.left.Key == nil {
			return nil
		}
		return find(key, node.left)
	}
}

func getRoot(node *Node) *Node {
	if node.parent != nil {
		return getRoot(node.parent)
	}
	return node
}

//return need balancing
func insertion(ins *Node, root *Node) bool {
	if root.Key.Equal(ins.Key) {
		root.Value = ins.Value
		return false
	}
	if root.Key.Compare(ins.Key) {
		if root.right.Key == nil {
			root.right = ins
			ins.parent = root
		} else {
			return insertion(ins, root.right)
		}
	} else {
		if root.left.Key == nil {
			root.left = ins
			ins.parent = root
		} else {
			return insertion(ins, root.left)
		}
	}
	return true
}

func balancing(node *Node) {
	if node.parent == nil || node.parent.parent == nil {
		return
	}
	grandPa := node.parent.parent
	father := node.parent
	var uncle *Node
	if grandPa.left == father {
		uncle = grandPa.right
	} else {
		uncle = grandPa.left
	}
	if !father.isBlack && uncle.isBlack {
		switch {
		case grandPa.left == father && father.right == node:
			turnLeft(node)
			balancing(node.left)
			return
		case grandPa.right == father && father.left == node:
			turnRight(node)
			balancing(node.right)
			return
		case grandPa.right == father && father.right == node && !node.isBlack && !father.isBlack:
			turnLeft(father)
			grandPa.isBlack = false
			father.isBlack = true
			return
		case grandPa.left == father && father.left == node && !node.isBlack && !father.isBlack:
			turnRight(father)
			grandPa.isBlack = false
			father.isBlack = true
			return
		}
	}
	if !uncle.isBlack && !father.isBlack {
		grandPa.isBlack = false
		father.isBlack = true
		uncle.isBlack = true
		balancing(grandPa)
	}

}

func turnLeft(node *Node) {
	father := node.parent
	node.parent = father.parent
	father.parent = node
	father.right = node.left
	node.left.parent = father
	node.left = father
	if node.parent != nil {
		if node.parent.left == father {
			node.parent.left = node
		} else if node.parent.right == father {
			node.parent.right = node
		}
	}
}

func turnRight(node *Node) {
	father := node.parent
	node.parent = father.parent
	father.parent = node
	father.left = node.right
	node.right.parent = father
	node.right = father
	if node.parent != nil {
		if node.parent.left == father {
			node.parent.left = node
		} else if node.parent.right == father {
			node.parent.right = node
		}
	}
}

func checkInserting(ins *Node) error {
	if ins.Key == nil {
		return errors.New("inserting node must have key")
	}
	return nil
}

//delete node, return parent if need to be balanced
func deleting(node *Node) *Node {
	if node.isBlack {
		if node.left.Key != nil {
			node.Key = node.left.Key
			node.left.Key = nil
			node.left.left = nil
			node.left.right = nil
			if node.left.isBlack {
				return node.left
			} else {
				node.left.isBlack = true
				return nil
			}
		} else if node.right.Key != nil {
			node.Key = node.right.Key
			node.right.Key = nil
			node.right.left = nil
			node.right.right = nil
			if node.right.isBlack {
				return node.right
			} else {
				node.right.isBlack = true
				return nil
			}
		} else {
			node.Key = nil
			node.left = nil
			node.right = nil
			return node
		}
	} else {
		node.Key = nil
		node.left = nil
		node.right = nil
		node.isBlack = true
	}
	return nil
}

func findMaxLeft(node *Node) *Node {
	if node.left.Key == nil {
		return node
	}
	maxL := node.left
	for maxL.right.Key != nil {
		maxL = maxL.right
	}
	return maxL
}

//delete cases match the https://en.wikipedia.org/wiki/Red–black_tree
func delBalancing(bal *Node) {
	if bal != nil && bal.isBlack && bal.parent != nil {
		father := bal.parent
		var isLeft bool
		var brother *Node
		if father.left == bal {
			brother = father.right
			isLeft = true
		} else {
			brother = father.left
		}
		lNephew := brother.left
		rNephew := brother.right

		//case1
		if brother.isBlack && father.isBlack && lNephew.isBlack && rNephew.isBlack {
			brother.isBlack = false
			delBalancing(father)
			return
		}
		//case4
		if !father.isBlack && lNephew.isBlack && rNephew.isBlack && brother.isBlack {
			brother.isBlack = false
			father.isBlack = true
			return
		}
		//case3
		if !brother.isBlack && lNephew.Key != nil && rNephew.Key != nil && lNephew.isBlack && rNephew.isBlack {
			if isLeft {
				turnLeft(brother)
			} else {
				turnRight(brother)
			}
			father.isBlack = false
			brother.isBlack = true
			delBalancing(bal)
			return
		}
		var distNephew *Node
		var closeNephew *Node
		if isLeft {
			distNephew = rNephew
			closeNephew = lNephew
		} else {
			distNephew = lNephew
			closeNephew = rNephew
		}

		//case5
		if brother.isBlack && closeNephew != nil && !closeNephew.isBlack && distNephew.isBlack {
			if isLeft {
				turnRight(closeNephew)
			} else {
				turnLeft(closeNephew)
			}
			closeNephew.isBlack = true
			brother.isBlack = false
			delBalancing(bal)
			return
		}

		//case6
		if brother.isBlack && distNephew != nil && !distNephew.isBlack {
			if isLeft {
				turnLeft(brother)
			} else {
				turnRight(brother)
			}
			distNephew.isBlack = true
			brother.isBlack = father.isBlack
			father.isBlack = true
		}
	}
}
