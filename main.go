package main

import (
	bst "github.com/weezy20/tree-go/bst"
)

func main() {
	tree := bst.New[int](3, 5, 7)
	tree.Insert(2)
	// tree.PrintTree()
}
