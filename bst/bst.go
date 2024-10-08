/// Binary Search Tree (BST)

package bst

import (
	"cmp"
	"errors"
	"fmt"
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

type NodeType int

const (
	Left NodeType = iota
	Right
	Root
)

// Returns which child of the parent this node is
func (n *Node[T]) ChildType() NodeType {
	if n.Parent == nil {
		return Root
	} else if n.Parent.Left == n {
		return Left
	} else if n.Parent.Right == n {
		return Right
	}
	return Root // Default return to handle all cases
}

// Size returns the size of a Node
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

// Get non null children of a Node
func (n *Node[T]) Children() []*Node[T] {
	var children []*Node[T]
	if n.Left != nil {
		children = append(children, n.Left)
	}
	if n.Right != nil {
		children = append(children, n.Right)
	}
	return []*Node[T]{n.Left, n.Right}
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
		parent, err := t.findAvailableNode(&key)
		if err != nil {
			fmt.Println(err)
			return
		}
		if insertChild(parent, &key) {
			fmt.Printf("key %v inserted successfully\n", key)
		}
	}
}

// findAvailableNode finds a node in the tree where a new key can be inserted
// In case keys are already present, returns the node and an associated error
func (t *Tree[T]) findAvailableNode(key *T) (*Node[T], error) {
	current := t.Root // non-nil has already been checked in caller
	for current != nil && !current.isPartialLeaf() {
		if cmp.Less(*key, *current.Key) {
			current = current.Left
		} else if cmp.Less(*current.Key, *key) {
			current = current.Right
		} else {
			return current, errors.New("key already exists in the tree")
		}
	}
	return current, nil
}

// Assumes n is non-nil
// It simply inserts a key as a child of node n.
// n is expected to be at least a partial leaf checked above
// Node is checked to be non-nil by the functions above
func insertChild[T Item](parent *Node[T], key *T) bool {
	if parent == nil {
		panic("Expected parent node to be non-nil")
	}
	if parent.Left == nil && cmp.Less(*key, *parent.Key) {
		parent.Left = &Node[T]{Key: key, Parent: parent}
		return true
	} else if parent.Right == nil && cmp.Less(*parent.Key, *key) {
		parent.Right = &Node[T]{Key: key, Parent: parent}
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

// Delete a key from the tree
func (t *Tree[T]) Delete(key T) error {
	if *t.Root.Key == key {
		// Root is the only node in the tree
		if t.Root.isLeaf() {
			t.Root = nil
			return nil
		} else if t.Root.isPartialLeaf() {
			// Promote the only child as the new root
			child := t.Root.Children()[0]
			t.Root = child
			return nil
		} else {
			// Case 3: Root has two children
			// Find the minimum value in the right subtree
			min := findMin(t.Root.Right)
			t.Root.Key = min.Key
			// Delete the minimum value node
			delete(t.Root.Right, min.Key)
			return nil
		}
	}
	err := delete(t.Root, &key)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

// Delete a key from subtree
// If the node is a leaf node, then we simply remove it
// If the node is a partial leaf node, we promote the only child as the new node
// TODO : for two children
func delete[T Item](subtree *Node[T], key *T) error {
	if subtree == nil {
		return errors.New("key to remove not found in the tree")
	}
	if *subtree.Key == *key {
		// Case 1: Subtree is a leaf node
		if subtree.isLeaf() {
			var parent = subtree.Parent
			switch subtree.ChildType() {
			case Left:
				parent.Left = nil
			case Right:
				parent.Right = nil
			case Root:
				// Dead code, but still shown for completeness
				panic("root should be checked before calling delete")
			}
			zeroize(subtree)
			return nil
		}
		// Case 2: Subtree is a partial leaf node
		if subtree.isPartialLeaf() {
			var parent = subtree.Parent
			var child = subtree.Children()[0]
			switch subtree.ChildType() {
			case Left:
				parent.Left = child
			case Right:
				parent.Right = child
			}
			zeroize(subtree)
			return nil
		}
		// Case 3: Subtree has two children
		panic("Not implemented yet")
	} else {
		switch cmp.Less(*key, *subtree.Key) {
		case true:
			return delete(subtree.Left, key)
		case false:
			return delete(subtree.Right, key)
		}
	}
	return nil
}

// Zero a node
func zeroize[T Item](node *Node[T]) {
	node.Left = nil
	node.Right = nil
	node.Key = nil
	node.Parent = nil
}

func findMin[T Item](node *Node[T]) *Node[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
