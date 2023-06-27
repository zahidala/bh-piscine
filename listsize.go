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

// func ListPushFront(l *List, data interface{}) {
// 	newNode := &NodeL{Data: data}

// 	if l.Head == nil {
// 		l.Head = newNode
// 		l.Tail = newNode
// 		return
// 	}

// 	newNode.Next = l.Head
// 	l.Head = newNode
// }

// func ListSize(l *List) int {
// 	it := l.Head
// 	count := 0
// 	for it != nil {
// 		count = count + 1
// 		it = it.Next
// 	}
// 	return count
// }

// func main() {
// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "2")
// 	ListPushFront(link, "you")
// 	ListPushFront(link, "man")

// 	fmt.Println(ListSize(link))
// }
