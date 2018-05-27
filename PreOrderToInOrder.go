package main

import (
	"fmt"
	"os"
)

type Node struct {
	Data int
	Left *Node
	Right *Node
}


func print(curr *Node) {

	if curr == nil {
		return
	}

	fmt.Println(curr.Data)

	print(curr.Left)

	print(curr.Right)

}

func newNode(data int) Node {

	node := Node{}
	node.Data = data

	return node
}



type Stack []*Node
func (s *Stack) pushStack(n *Node)  {
	*s = append(*s, n)
}
func(s *Stack) popStack() *Node {
	length := len(*s)
	p := (*s)[length-1]
	*s = (*s)[:length-1]
	return p
}
func(s Stack) stackIsEmpty() bool{
	return len(s) == 0
}


func reverse(node *Node) intArray  {

	var stack Stack = Stack{}
	var inOrder intArray

	for node != nil || !stack.stackIsEmpty() {

		if node != nil {
			stack.pushStack(node)
			node = node.Left

		} else {
			node = stack.popStack()
			inOrder.push(node.Data)
			node = node.Right
		}

	}

	return inOrder

}



func push(node *Node, data int) {

	if node.Left == nil {
		node.Left = &Node{data, nil, nil}
	} else if node.Right == nil {
		node.Right = &Node{data, nil, nil}
	} else {
		push(node.Left, data)
	}

}

type intArray []int
func (a *intArray) push(value int) {
	*a = append(*a, value)
}

func (a intArray) print() {
	var i int = 0;
	for i=0; i< len(a); i++ {
		fmt.Println(a[i])
	}
}

func main() {


	for {

		os.Clearenv()

		fmt.Print("Input N:")
		var n int
		fmt.Scan(&n)

		fmt.Println("N=", n)

		if n >=1 && n <=1000 {

			var root *Node

			var i int
			var n1 = int(n)
			for i = 1; i <= n1; i++ {

				if i ==1 {
					root = &Node {i, nil,nil}
				} else {
					push(root, i)
				}
			}

			print(root)

			fmt.Println("Output:")
			var inOrder intArray = reverse(root)
			inOrder.print()

			break

		} else {
			fmt.Println("1 <= N <= 1000")
		}

	}





}
