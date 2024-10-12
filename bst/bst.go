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

// Every node, except Root, is a Child node, either Left/Right of it's parent node
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

// Node is a partial leaf if it has one child
func (n *Node[T]) isPartialLeaf() bool {
	return (n.Left == nil && n.Right != nil) || (n.Left != nil && n.Right == nil)
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
	return children
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
		if !insertChild(parent, &key) {
			fmt.Printf("Cannot insert key %v\n", key)
		}
	}
}

// findAvailableNode finds a node in the tree where a new key can be inserted
// In case keys are already present, returns the node and an associated error
func (t *Tree[T]) findAvailableNode(key *T) (*Node[T], error) {
	current := t.Root // Caller ensures this will never be nil
	for current != nil {
		if cmp.Less(*key, *current.Key) {
			if current.Left == nil {
				return current, nil
			}
			current = current.Left
		} else if cmp.Less(*current.Key, *key) {
			if current.Right == nil {
				return current, nil
			}
			current = current.Right
		} else {
			return current, errors.New("key already exists in the tree")
		}
	}
	if current == nil {

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
	node := search(t.Root, &key) // Will be nil if t.Root is nil
	if node == nil {
		return fmt.Errorf("key %v not found in the tree", key)
	}
	err := delete(t, node)
	if err != nil {
		return err
	}
	return nil
}

// Delete a key (z) from tree (t)
// If the node is a leaf node, then we simply remove it
// If the node is a partial leaf node, we promote the only child as the new node
// For the third case, where the node has two children:
// - Find the node's in-order successor (the smallest node in the right subtree).
// - Replace the value of the node to be deleted with the value of the in-order successor.
// - Recursively delete the in-order successor, which will now be a simpler case
// (either a leaf node or a node with one child).
func delete[T Item](t *Tree[T], z *Node[T]) error {
	// Case 1: Node (z) is a leaf node
	if z.isLeaf() {
		var parent = z.Parent // maybe nil if z is root node itself
		switch z.ChildType() {
		case Left:
			parent.Left = nil
		case Right:
			parent.Right = nil
		case Root:
			t.Root = nil
		}
		zeroize(z)
		return nil
	}
	// Case 2: Subtree is a partial leaf node
	if z.isPartialLeaf() {
		var parent = z.Parent
		var child = z.Children()[0] // Guaranteed to have only one child
		switch z.ChildType() {
		case Left:
			parent.Left = child
		case Right:
			parent.Right = child
		case Root:
			t.Root = child
		}
		zeroize(z)
		return nil
	}
	// Case 3: Subtree has two children, check unnecessary but shown for clarity
	if z.Left != nil && z.Right != nil {
		// Note: the same set of values may be represented as different binary-search trees
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
func findMax[T Item](node *Node[T]) *Node[T] {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// PrintTree prints the tree structure in the CLI using '\' and '/' for showing links
