package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	Value int // interface{}
}

// FrontTraverse 前序遍历，根-左-右
func FrontTraverse(t *BinaryTree) {
	if t == nil {
		return
	}
	fmt.Print(t.Value, " ")
	FrontTraverse(t.Left)
	FrontTraverse(t.Right)
}

// NewBinaryTree 返回一个BinaryTree
func NewBinaryTree(n int) *BinaryTree {
	var t *BinaryTree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		value := rand.Intn(2 * n)
		t = insert(t, value)
	}
	return t
}

func insert(t *BinaryTree, value int) *BinaryTree {
	// 创建根节点
	if t == nil {
		return &BinaryTree{nil, nil, value}
	}
	if value == t.Value {
		return t
	}
	if value < t.Value {
		t.Left = insert(t.Left, value)
		return t
	}
	t.Right = insert(t.Right, value)
	return t
}

func main() {
	tree := NewBinaryTree(10)
	fmt.Println("The root of the tree is ", tree.Value)
	FrontTraverse(tree)
	fmt.Println()
}
