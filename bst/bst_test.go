package bst

import (
	"testing"
)

func TestBST(t *testing.T) {
	// Test creating a new BST with no arguments
	tree := New[int]()
	if tree.Root != nil {
		t.Errorf("Expected Root to be nil, got %v", tree.Root)
	}

	// Test creating a new BST with multiple arguments
	tree = New(10, 5, 15, 3, 7, 12, 18)
	if tree.Root == nil {
		t.Errorf("Expected Root to be non-nil")
	}

	// Test inserting into the BST
	tree.Insert(6)
	if tree.Root.Left.Right.Left.Key == nil || *tree.Root.Left.Right.Left.Key != 6 {
		t.Errorf("Expected to find key 6 in the tree")
	}

	// Test in-order traversal
	expectedOrder := []int{3, 5, 6, 7, 10, 12, 15, 18}
	inOrderResult := tree.InOrder()
	for i, v := range expectedOrder {
		if inOrderResult[i] != v {
			t.Errorf("Expected in-order result to be %v, got %v", expectedOrder, inOrderResult)
		}
	}

	// Test isLeaf method
	if !tree.Root.Left.Left.isLeaf() {
		t.Errorf("Expected node with key 3 to be a leaf")
	}

	// Test isPartialLeaf method
	if !tree.Root.Left.isPartialLeaf() {
		t.Errorf("Expected node with key 5 to be a partial leaf")
	}

	// Test Size method
	if tree.Root.Size() == 0 {
		t.Errorf("Expected non-zero size for the root node")
	}
}
