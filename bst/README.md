# Binary Search Tree (BST)

A Binary Search Tree (BST) is a classical tree data structure that facilitates fast lookup, addition, and deletion of items. Each node in a BST contains a key and two distinguished sub-trees, commonly referred to as the left and right sub-trees. The left sub-tree contains nodes with keys less than the node's key, and the right sub-tree contains nodes with keys greater than the node's key.

## Features
- **Efficient Search**: O(log n) time complexity for search operations.
- **Dynamic Insertion and Deletion**: Allows for dynamic data management.
- **Ordered Structure**: Maintains elements in a sorted order.
- **Easy to understand**: The simplest tree data structure out there

## Disadvantages
- Not the most practical tree data structure in terms of minimizing IOPs 
- Cache misses due to following pointers of heap allocated data


## Usage

A `Tree` is a struct containing a pointer to a `Root` node which is of type `Node[T]` and is 32 bytes
It doesn't matter if `T` is heap allocated as long as it follows `cmp.Ordered` and `comparable` interfaces.

For each Node we visit, we have to follow it's parent's pointer, and then the key's pointer
so that is 2 pointers per node, 3 if we have to descend further down one of the child pointer of a node.

Below are the public functions exposed in `bst.go` and their usage:

### Create a BST

```go
import "github.com/weezy20/bst"

// Creates an empty tree with root node = nil
var tree = bst.New() 

// Creates a tree with one root node with key 1
var tree = bst.New(1)

// Creates a tree with one root node with key 4 
// And subsequently inserts the rest of the arguments into the tree
var tree = bst.New(4, 5, 12, 24, 44) 
```

### Insert
Inserts a new key into the BST.

```go
func (t *Tree[T]) Insert(key T)
```

**Example:**
```go
var tree = bst.New[int]()
tree.Insert(4)
```

### Search
Searches for a key in the BST.

```go
func (t *Tree[T]) Search(key T) *Node[T]
```

**Example:**
```go
tree = bst.New[string]("Hello")
node := tree.Search("Hello") // Returns root node
```

### Delete
Deletes a key from the BST.

```go
func (t *Tree[T]) Delete(key T)
```

**Example:**
```go
tree.Delete(10)
```

### InOrder Traversal
Performs an in-order traversal of the BST.

**Example:**
```go
keys := bst.InOrder() // returns keys in sorted order
```

## Testing

Elaborate usage examples can be found in `bst_test.go`.
```sh
go test ./bst
```



