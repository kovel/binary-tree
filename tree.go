package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"bufio"
	"os"
	"strconv"
	"math"
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

func (t *Tree)Find(key int) *Node {
	if t.Root == nil {
		return nil
	}

	current := t.Root
	for current.Key != key {
		if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}

		if current == nil {
			return nil
		}
	}

	return current
}

func (t *Tree)Delete(key int) bool {
	isLeftKey := false
	current := t.Root;
	parent := current
	for current.Key != key {
		parent = current
		if key < current.Key {
			isLeftKey = true;
			current = current.Left
		} else {
			isLeftKey = false
			current = current.Right
		}

		if current == nil {
			return false
		}
	}

	if current.Left == nil && current.Right == nil {
		if current == t.Root {
			t.Root = nil
		} else if isLeftKey {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if current.Left != nil && current.Right == nil {
		if current == t.Root {
			t.Root = current.Left
		} else if isLeftKey {
			parent.Left = current.Left
			current.Left = nil
		} else {
			parent.Right = current.Left
			current.Left = nil
		}
	} else if current.Left == nil && current.Right != nil {
		if current == t.Root {
			t.Root = current.Right
		} else if isLeftKey {
			parent.Left = current.Right
			current.Right = nil
		} else {
			parent.Right = current.Right
			current.Right = nil
		}
	} else {
		successor := t.getSuccessor(current)
		if current == t.Root {
			t.Root = successor
		} else if isLeftKey {
			parent.Left = successor
		} else {
			parent.Right = successor
		}
		successor.Left = current.Left
	}

	return true
}

func (t *Tree)getSuccessor(node *Node) *Node {
	successorParent := node
	successor := node.Right

	for successor.Left != nil {
		successorParent = successor
		successor = successor.Left
	}

	if successor != node.Right {
		successorParent.Left = successor.Right
		successor.Right = node.Right
	}

	return successor
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
			fmt.Print(strings.Repeat(" ", int(math.Abs(float64(n*2-2)))))
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

	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("Enter first letter of operation to proceed(delete, find, insert, traverse, print): ")
		operation, _, _ := reader.ReadLine()
		switch operation[0] {
		case "f":
			fmt.Print("Enter key to find: ")
			line, _, _ := reader.ReadLine()
			key, _ := strconv.Atoi(string(line))
			found := tree.Find(key)
			if found != nil {
				fmt.Println(found)
			} else {
				fmt.Printf("Not with key %d not found\n", key)
			}

			break
		case "d":
			fmt.Print("Enter key to delete: ")
			line, _, _ := reader.ReadLine()
			key, _ := strconv.Atoi(string(line))
			tree.Delete(key)
			break
		case "i":
			fmt.Print("Enter key to insert: ")
			line, _, _ := reader.ReadLine()
			key, _ := strconv.Atoi(string(line))
			tree.Insert(&Node{Key: key, Value: fmt.Sprintf("value #%d", key)})
			break
		case "t":
			tree.Traverse()
		case "p":
			tree.PrintTree()
			break
		}
	}
}
