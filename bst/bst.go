/// Binary Search Tree (BST)

package bst

import (
	"cmp"
	"unsafe"
)

type Item interface {
	cmp.Ordered // <, <=, >, >= operators
	comparable  // == and != operators
}

type Tree[T Item] struct {
	Root *Node[T]
}

// Heap allocated node
type Node[T Item] struct {
	// Key of a Node
	Key *T
	// Left Child of a Node
	Left *Node[T]
	// Right Child of a Node
	Right *Node[T]
	// Parent of a Node
	Parent *Node[T]
}

// Size returns the number of nodes in the tree
func (n *Node[T]) Size() uintptr {
	return unsafe.Sizeof(*n)
}

// Create a new BST. In case multiple arguments are passed, the first one is considered the root
// and subsequent ones are inserted into the tree.
// In case no arguments are passed, a tree with nil Root is returned
func New[T Item](args ...T) Tree[T] {
	var root *Node[T]
	tree := Tree[T]{Root: root}
	if len(args) > 0 {
		for _, arg := range args {
			tree.Insert(arg)
		}
	}
	return tree
}

func (t *Tree[T]) Insert(key T) {
	if t.Root == nil {
		t.Root = &Node[T]{Key: &key}
		return
	}
	insertChild(t.Root, &key)
}

// This function doesn't check BST traversal constraints
// It simply inserts a key as a child of node `n` if it has space for children
// Returns `true` on successful insertion otherwise `false` for failure
func insertChild[T Item](n *Node[T], key *T) bool {
	// Primary condition we want to check
	if n == nil {
		return false
	}
	// In case of equal keys, we prefer the Left Child to hold the key
	if n.Left == nil && (cmp.Less(*key, *n.Key) || *key == *n.Key) {
		n.Left = &Node[T]{Key: key, Parent: n}
		return true
	}
	if n.Right == nil && cmp.Less(*n.Key, *key) {
		n.Right = &Node[T]{Key: key, Parent: n}
		return true
	}
	return false
}
