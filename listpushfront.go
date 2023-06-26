package piscine

// package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	newNode.Next = l.Head
	l.Tail = newNode
}

// func main() {

// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "man")
// 	ListPushFront(link, "how are you")

// 	it := link.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " ")
// 		it = it.Next
// 	}
// 	fmt.Println()
// }
