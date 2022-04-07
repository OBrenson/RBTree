# RBTree
Red Black Tree on Golang

Simple realisation of red black tree with simple api.

    go get -v github.com/OBrenson/RBTree

Examples of usages

Implementation of Key interface:

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

Create tree with root node:

    tree := CreateTree(NodeVal(0))
    
Insertion(second param is node`s value, it can be any struct)
    
    err := tree.Insert(NodeVal(rand.Int31n(100)), nil)
    
Remove:

    nfErr := tree.Remove(13) //remove key equal 13
    
More examples you can find in tree_test.go file.