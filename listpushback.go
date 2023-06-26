package piscine

// package main

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

// func main() {

// 	link := &List{}

// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, "man")
// 	ListPushBack(link, "how are you")

// 	for link.Head != nil {
// 		fmt.Println(link.Head.Data)
// 		link.Head = link.Head.Next
// 	}
// }
