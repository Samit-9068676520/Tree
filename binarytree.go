package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type Node struct {
	value       int
	left, right *Node
}
type Tree struct {
	root *Node
}

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}
func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}
func preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	preOrder(n.left)
	preOrder(n.right)
}
func inOrder(n *Node) {
	if n == nil {
		return
	}
	inOrder(n.left)
	fmt.Print(n.value, " ")
	inOrder(n.right)
}
func postOrder(n *Node) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	fmt.Print(n.value, " ")
}
func breadthfirstSearch(n *Node) {
	q := queue.New()
	if n != nil {
		q.Enqueue(n)
	}
	for q.Len() != 0 {
		temp := q.Dequeue().(*Node)
		fmt.Print(temp.value, " ")
		if temp.left != nil {
			q.Enqueue(temp.left)
		}
		if temp.right != nil {
			q.Enqueue(temp.right)
		}
	}
	fmt.Println()
}
func defthfirstSearch(n *Node) {
	s := stack.New()
	if n != nil {
		s.Push(n)
	}
	for s.Len() != 0 {
		temp := s.Pop().(*Node)
		fmt.Print(temp.value, " ")
		if temp.left != nil {
			s.Push(temp.left)
		}
		if temp.right != nil {
			s.Push(temp.right)
		}
	}
	fmt.Println()
}
func NthpreOrder(n *Node, index int, counter *int) {
	if n == nil {
		return
	}
	(*counter)++
	if (*counter) == index {
		fmt.Print(n.value, " ")
		return
	}
	NthpreOrder(n.left, index, counter)
	NthpreOrder(n.right, index, counter)
}
func copyTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.left = copyTree(curr.left)
		temp.right = copyTree(curr.right)
		return temp
	}
	return nil
}
func copyMirrorTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.right = copyMirrorTree(curr.left)
		temp.left = copyMirrorTree(curr.right)
		return temp
	}
	return nil
}
func numberofleafNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	}
	return numberofleafNodes(curr.left) + numberofleafNodes(curr.right)
}
func numberofNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	return 1 + numberofNodes(curr.left) + numberofNodes(curr.right)
}
func identical(n1 *Node, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil {
		return false
	}
	leftsubtree := identical(n1.left, n2.left)
	rightsubtree := identical(n1.right, n2.right)
	return (n1.value == n2.value) && leftsubtree && rightsubtree
}
func sumAllBT(curr *Node) int {
	var sum, leftsum, rightsum int
	if curr == nil {
		return 0
	}
	rightsum = sumAllBT(curr.right)
	leftsum = sumAllBT(curr.left)
	sum = curr.value + leftsum + rightsum
	return sum
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	t1 := LevelOrderBinaryTree(arr)
	preOrder(t1.root)
	fmt.Println()
	inOrder(t1.root)
	fmt.Println()
	postOrder(t1.root)
	fmt.Println("=================")
	fmt.Println("Breadth First Search")
	breadthfirstSearch(t1.root)
	fmt.Println("================")
	fmt.Println("Defth First Search")
	defthfirstSearch(t1.root)
	fmt.Println("===============")
	fmt.Println("Nth Pre Order")
	var counter int
	NthpreOrder(t1.root, 3, &counter)
	fmt.Println("==================")
	fmt.Println("CopyTree")
	t2 := new(Tree)
	t2.root = copyTree(t1.root)
	fmt.Println("Pre Order of Copy Tree")
	preOrder(t2.root)
	fmt.Println("==============")
	fmt.Println("CopyMirrorTree")
	t3 := new(Tree)
	t3.root = copyMirrorTree(t1.root)
	fmt.Println("Pre Order of CopyMirrorTree")
	preOrder(t3.root)
	fmt.Println("=============")
	fmt.Printf("No of nodes in a Binary Tree: %v", numberofNodes(t1.root))
	fmt.Println()
	fmt.Println("=============")
	fmt.Printf("No of leaf nodes in a Binary Tree: %v", numberofleafNodes(t1.root))
	fmt.Println()
	fmt.Print("============")
	fmt.Printf("Check identical of t1 and t2: %v", identical(t1.root, t2.root))
	fmt.Println()
	fmt.Printf("Check identical of t2 and t3: %v", identical(t2.root, t3.root))
	fmt.Println()
	fmt.Printf("Check identical of t3 and t1: %v", identical(t1.root, t3.root))
	fmt.Println("===============")
	fmt.Println("Sum of All Node")
	fmt.Printf("Sum of All Nodes in t1: %v", sumAllBT(t1.root))
}
