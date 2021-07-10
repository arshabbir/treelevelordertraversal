package main

import (
	"log"
	"time"
)

type Node struct {
	left  *Node
	right *Node
	data  int
}

//Implement tree building functionlity

func (n *Node) Insert(data int) {

	if data == n.data {
		log.Println("Data already present")
		return
	}

	if data < n.data {
		if n.left == nil {
			t := new(Node)
			t.data = data
			t.left = nil
			t.right = nil
			n.left = t

		} else {
			n.left.Insert(data)
		}
		return
	}

	if data > n.data {
		if n.right == nil {
			t := new(Node)
			t.data = data
			t.left = nil
			t.right = nil
			n.right = t

		} else {
			n.right.Insert(data)
		}
		return
	}

}

func (n *Node) Preorder() {
	if n == nil {
		return
	}
	log.Println(n.data)
	n.left.Preorder()
	n.right.Preorder()

}

func (n *Node) Postorder() {
	if n == nil {
		return
	}

	n.left.Postorder()
	n.right.Postorder()
	log.Println(n.data)

}

type Que struct {
	front int
	rear  int
	size  int
	que   []*Node
}

//Implement the Que functionality

func NewQue() *Que {
	return &Que{front: 0, rear: 0, que: []*Node{}}
}

func (q *Que) enque(n *Node) {
	q.que = append(q.que, n)
	q.rear++
	q.size++

}
func (q *Que) deque() *Node {

	if q.size == 0 {
		return nil
	}
	n := q.que[q.front]

	q.que = q.que[q.front+1:]
	q.size--

	//q.front = 0
	return n

}

func main() {

	q := NewQue()

	//Build Tree
	tree := &Node{data: 150}

	tree.Insert(50)
	tree.Insert(20)
	tree.Insert(100)
	tree.Insert(250)
	tree.Insert(230)
	tree.Insert(280)
	tree.Insert(75)
	tree.Insert(65)
	tree.Insert(120)

	log.Println("Preorder traversal ")
	tree.Preorder()

	log.Println("PostOrder traversal")
	tree.Postorder()

	log.Println("Level order traversal")

	q.enque(tree)
	tree.LevelOrderTraversal(q)

}

func (n *Node) LevelOrderTraversal(q *Que) {

	//Deque the node
	//for {
	if node := q.deque(); node != nil {
		log.Println(node.data)
		if node.left != nil {
			q.enque(node.left)

		}
		if node.right != nil {
			q.enque(node.right)
		}

	} else {
		log.Println("Deque returned nil")
		return
	}

	time.Sleep(time.Second)
	//}
	n.LevelOrderTraversal(q)

}
