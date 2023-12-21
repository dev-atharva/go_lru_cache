package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const size = 5

type Node struct {
	val   string
	left  *Node
	right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.right = tail
	tail.left = head

	return Queue{
		Head: head,
		Tail: tail,
	}
}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)

	} else {
		node = &Node{
			val: str,
		}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove %s\n", n.val)
	left := n.left
	right := n.right

	left.right = right
	right.left = left

	c.Queue.Length -= 1
	delete(c.Hash, n.val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add:%s\n", n.val)
	tmp := c.Queue.Head.right

	c.Queue.Head.right = n
	n.left = c.Queue.Head
	n.right = tmp
	tmp.left = n

	c.Queue.Length++
	if c.Queue.Length > size {
		c.Remove(c.Queue.Tail.left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Start cache")
	cache := NewCache()

	for {
		fmt.Print("Enter a word (or 'exit' to quit): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		}

		cache.Check(input)
		cache.Display()
	}
}
