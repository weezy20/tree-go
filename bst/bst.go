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

// Node is a Leaf if it has no children
func (n *Node[T]) isLeaf() bool {
	return n.Left == nil && n.Right == nil
}

// Node is a partial leaf if it has at least one child
func (n *Node[T]) isPartialLeaf() bool {
	if n.isLeaf() {
		return true
	} else {
		return (n.Left == nil && n.Right != nil) || (n.Left != nil && n.Right == nil)
	}
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
	} else {
		node := t.findAvailableNode(&key)
		insertChild(node, &key)
	}
}

// findAvailableNode finds a node in the tree where a new key can be inserted
func (t *Tree[T]) findAvailableNode(key *T) *Node[T] {
	current := t.Root
	for current != nil && !current.isPartialLeaf() {
		if cmp.Less(*key, *current.Key) || *key == *current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return current
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

// InOrder traversal of the tree. Returns a slice of keys in sorted order
func (t *Tree[T]) InOrder() []T {
	var result []T
	inOrder(t.Root, &result)
	return result
}

func inOrder[T Item](n *Node[T], result *[]T) {
	if n == nil {
		return
	}
	inOrder(n.Left, result)
	*result = append(*result, *n.Key)
	inOrder(n.Right, result)
}

// Search returns the node ptr with the key if found, nil otherwise
func (t *Tree[T]) Search(key T) *Node[T] {
	return search(t.Root, &key)
}

func search[T Item](n *Node[T], key *T) *Node[T] {
	if n == nil || *n.Key == *key {
		return n
	}
	if cmp.Less(*key, *n.Key) {
		return search(n.Left, key)
	} else {
		return search(n.Right, key)
	}
}
