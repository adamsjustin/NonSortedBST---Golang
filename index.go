package main

import (
	"fmt"
	"math"
	"sync"
)

//Node ...
type Node struct {
	value int
	left  *Node //left
	right *Node //right
}

//NewNode ...
type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

var currentHeight int = 0

func (bst *BinarySearchTree) Insert(value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func insertNode(curNode, newNode *Node) {
	if newNode.value < curNode.value {
		if curNode.left == nil {
			curNode.left = newNode
		} else {
			insertNode(curNode.left, newNode)
		}
	}

	if newNode.value > curNode.value {
		if curNode.right == nil {
			curNode.right = newNode
		} else {
			insertNode(curNode.right, newNode)
		}
	}
}

func treeHeight(curNode *Node, curHeight int) int {
	if curNode == nil {
		return curHeight
	}
	leftHeight := treeHeight(curNode.left, (curHeight + 1))
	rightHeight := treeHeight(curNode.right, (curHeight + 1))
	return int(math.Max(float64(leftHeight), float64(rightHeight)))
}

func searchTree(curNode *Node, value int) bool {
	if curNode == nil {
		return false
	}
	if curNode.value == value {
		return true
	} else if value < curNode.value {
		return searchTree(curNode.left, value)
	} else if value > curNode.value {
		return searchTree(curNode.right, value)
	}
	return false
}

var bst BinarySearchTree

func fillTree(bst *BinarySearchTree) {
	bst.Insert(5)
	bst.Insert(4)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(9)
	bst.Insert(11)
	bst.Insert(12)
}

func main() {
	fillTree(&bst)
	fmt.Println(treeHeight(bst.root, currentHeight))
	fmt.Println(searchTree(bst.root, 13))
}
