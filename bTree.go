package tree

// binary tree node
type bTree struct {
	key         int
	left, right *bTree
}

// traverse the tree in order
// exec the f function received in each node
func indorder(node *bTree, f func(int)) {
	if node != nil {
		indorder(node.left, f)
		f(node.key)
		indorder(node.right, f)
	}
}

// traverse the tree in pre-order
// exec the f function received in each node
func preorder(node *bTree, f func(int)) {
	if node != nil {
		f(node.key)
		preorder(node.left, f)
		preorder(node.right, f)
	}
}

// traverse the tree in post-order
// exec the f function received in each node
func postorder(node *bTree, f func(int)) {
	if node != nil {
		postorder(node.left, f)
		postorder(node.right, f)
		f(node.key)
	}
}
