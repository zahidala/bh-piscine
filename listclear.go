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

// func PrintList(l *List) {
// 	link := l.Head
// 	for link != nil {
// 		fmt.Print(link.Data, " -> ")
// 		link = link.Next
// 	}
// 	fmt.Println(nil)
// }

// func ListPushBack(l *List, data interface{}) {
// 	newNode := &NodeL{Data: data}

// 	if l.Head == nil {
// 		l.Head = newNode
// 		return
// 	}

// 	curr := l.Head
// 	for curr.Next != nil {
// 		curr = curr.Next
// 	}

// 	curr.Next = newNode
// }

func ListClear(l *List) {
	l.Head = nil
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "I")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "something")
// 	ListPushBack(link, 2)

// 	fmt.Println("------list------")
// 	PrintList(link)
// 	ListClear(link)
// 	fmt.Println("------updated list------")
// 	PrintList(link)
// }
