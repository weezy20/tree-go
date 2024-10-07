package bst

import (
	"testing"
	"unsafe"
)

// Define a simple type for testing
type IntItem int

func (a IntItem) Less(b IntItem) bool {
	return a < b
}

func (a IntItem) Greater(b IntItem) bool {
	return a > b
}

func TestNewTree(t *testing.T) {
	// Test tree creation with no arguments
	tree := New[IntItem]()
	if tree.Root != nil {
		t.Errorf("Expected Root to be nil, got %v", tree.Root)
	}

	// Test tree creation with multiple elements
	tree = New[IntItem](5, 3, 7)
	if tree.Root == nil {
		t.Errorf("Expected Root to be non-nil, got nil")
	}
	if *tree.Root.Key != 5 {
		t.Errorf("Expected Root key to be 5, got %d", *tree.Root.Key)
	}
}

func TestInsert(t *testing.T) {
	// Test single insert
	tree := New[IntItem]()
	tree.Insert(10)
	if tree.Root == nil {
		t.Errorf("Expected Root to be non-nil after insertion, got nil")
	}
	if *tree.Root.Key != 10 {
		t.Errorf("Expected Root key to be 10, got %d", *tree.Root.Key)
	}

	// Test multiple inserts
	tree.Insert(5)
	tree.Insert(15)
	if *tree.Root.Left.Key != 5 {
		t.Errorf("Expected Left child to be 5, got %d", *tree.Root.Left.Key)
	}
	if *tree.Root.Right.Key != 15 {
		t.Errorf("Expected Right child to be 15, got %d", *tree.Root.Right.Key)
	}

	// Test inserting duplicate key
	// tree.Insert(5)
	// if tree.Root.Left.Left != nil {
	// 	t.Errorf("Expected Left-Left child to be nil (no duplicate insertion), got %v", tree.Root.Left.Left)
	// }
}

func TestInOrder(t *testing.T) {
	// Test in-order traversal
	tree := New[IntItem](10, 5, 15, 3, 7, 12, 17)
	expected := []IntItem{3, 5, 7, 10, 12, 15, 17}
	result := tree.InOrder()
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected in-order traversal value at index %d to be %d, got %d", i, expected[i], v)
		}
	}
}

func TestSearch(t *testing.T) {
	// Test search functionality
	tree := New[IntItem](10, 5, 15, 3, 7, 12, 17)

	// Test search for existing key
	node := tree.Search(7)
	if node == nil {
		t.Errorf("Expected to find key 7, but got nil")
	} else if *node.Key != 7 {
		t.Errorf("Expected node key to be 7, got %d", *node.Key)
	}

	// Test search for non-existent key
	node = tree.Search(20)
	if node != nil {
		t.Errorf("Expected to not find key 20, but got non-nil node")
	}
}

func TestIsLeaf(t *testing.T) {
	// Test isLeaf method
	tree := New[IntItem](10, 5, 15)
	if !tree.Root.Left.isLeaf() {
		t.Errorf("Expected node with key 5 to be a leaf")
	}
	if tree.Root.isLeaf() {
		t.Errorf("Expected root node with key 10 to not be a leaf")
	}
}

func TestIsPartialLeaf(t *testing.T) {
	// Test isPartialLeaf method
	tree := New[IntItem](10, 5)
	if !tree.Root.isPartialLeaf() {
		t.Errorf("Expected root node with key 10 to be a partial leaf")
	}
	if !tree.Root.Left.isLeaf() {
		t.Errorf("Expected left node with key 5 to be a leaf")
	}
}

func TestNodeSize(t *testing.T) {
	// Node struct is 4+4+4+4 = 32 bytes wide
	expectedSize := uintptr(32)

	var node *Node[IntItem]
	size := unsafe.Sizeof(*node)
	if size != expectedSize {
		t.Errorf("Expected node size to be %d bytes, got %d bytes", expectedSize, size)
	}

	var largeNode *Node[string]
	largeNodeSize := unsafe.Sizeof(*largeNode)
	if largeNodeSize != expectedSize {
		t.Errorf("Expected large node size to be %d bytes, got %d bytes", expectedSize, largeNodeSize)
	}
}
