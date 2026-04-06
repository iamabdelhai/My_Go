package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (b *BST) Insert(val int) {
	if b.Root == nil {
		b.Root = &Node{Val: val}
		return
	}
	current := b.Root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = &Node{Val: val}
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &Node{Val: val}
				return
			}
			current = current.Right
		}
	}
}

func (b *BST) Search(val int) bool {
	current := b.Root
	for current != nil {
		if val == current.Val {
			return true
		} else if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

func (b *BST) Delete(val int) {
	b.Root = deleteNode(b.Root, val)
}

func deleteNode(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if val < node.Val {
		node.Left = deleteNode(node.Left, val)
	} else if val > node.Val {
		node.Right = deleteNode(node.Right, val)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		successor := findMin(node.Right)
		node.Val = successor.Val
		node.Right = deleteNode(node.Right, successor.Val)
	}
	return node
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (b *BST) InOrder() []int {
	var result []int
	inOrder(b.Root, &result)
	return result
}

func inOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, result)
	*result = append(*result, node.Val)
	inOrder(node.Right, result)
}

func (b *BST) PreOrder() []int {
	var result []int
	preOrder(b.Root, &result)
	return result
}

func preOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preOrder(node.Left, result)
	preOrder(node.Right, result)
}

func (b *BST) Height() int {
	return height(b.Root)
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	leftH := height(node.Left)
	rightH := height(node.Right)
	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}

func main() {
	bst := &BST{}
	values := []int{50, 30, 70, 20, 40, 60, 80}
	fmt.Println("Inserting values:", values)
	for _, v := range values {
		bst.Insert(v)
	}
	fmt.Println("\n── Traversals ──")
	fmt.Println("In-Order  (sorted):", bst.InOrder())
	fmt.Println("Pre-Order         :", bst.PreOrder())
	fmt.Println("Tree Height       :", bst.Height())
	fmt.Println("\n── Search ──")
	fmt.Println("Search(40) →", bst.Search(40))
	fmt.Println("Search(99) →", bst.Search(99))
	fmt.Println("\n── Delete 30 ──")
	bst.Delete(30)
	fmt.Println("In-Order after delete:", bst.InOrder())
}