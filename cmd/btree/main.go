package main

import (
	"fmt"

	"mastering-go/internal/pkg/btree"
)

func main() {
	tree := btree.NewBinaryTree(10)
	fmt.Println("The root of the tree is ", tree.Value)
	btree.FrontTraverse(tree)
	fmt.Println()
}
