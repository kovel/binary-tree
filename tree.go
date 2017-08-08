package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Node struct {
	Left  *Node
	Right *Node

	Key   int
	Value string
}

func (n *Node) String() string {
	return fmt.Sprintf("{ %d: %s }", n.Key, n.Value)
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) IsEmpty() bool {
	return s.count <= 0
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(node *Node) {
	current := t.Root
	parent := current
	if current == nil {
		t.Root = node
		return
	}

	for true {
		parent = current
		if node.Key < current.Key {
			current = current.Left
			if current == nil {
				parent.Left = node
				return
			}
		} else {
			current = current.Right
			if current == nil {
				parent.Right = node
				return
			}
		}
	}
}

func (t *Tree) Traverse() {
	t.printNode(t.Root)
}

func (t *Tree) printNode(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node)
	t.printNode(node.Left)
	t.printNode(node.Right)
}

func (t *Tree) PrintTree() {
	n := 32
	isRowEmpty := false
	globalStack := &Stack{}
	globalStack.Push(t.Root)

	for !isRowEmpty {
		localStack := &Stack{}
		isRowEmpty = true

		fmt.Print(strings.Repeat(" ", n))
		for !globalStack.IsEmpty() {
			node := globalStack.Pop()
			if node != nil {
				fmt.Print(node.Key)
				localStack.Push(node.Left)
				localStack.Push(node.Right)

				if node.Left != nil || node.Right != nil {
					isRowEmpty = false
				}
			} else {
				localStack.Push(nil)
				localStack.Push(nil)
				fmt.Print("--")
			}
			fmt.Print(strings.Repeat(" ", n*2-2))
		}
		fmt.Println()

		for !localStack.IsEmpty() {
			globalStack.Push(localStack.Pop())
		}

		n /= 2
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("Filling tree with values...")
	tree := &Tree{}
	for i := 0; i < 10; i++ {
		value := rand.Intn(100)
		node := &Node{Key: value, Value: fmt.Sprintf("value #%d", value)}
		tree.Insert(node)
	}

	fmt.Println("\nTraversing...")
	tree.Traverse()

	fmt.Println("\nPrinting tree")
	tree.PrintTree()
}
