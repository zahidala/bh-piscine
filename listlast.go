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

func ListLast(l *List) interface{} {
	it := l.Head

	if it == nil {
		return nil
	}

	for it.Next != nil {
		it = it.Next
	}
	return it.Data
}

// func main() {
// 	link := &List{}
// 	link2 := &List{}

// 	ListPushBack(link, "three")
// 	ListPushBack(link, 3)
// 	ListPushBack(link, "1")

// 	fmt.Println(ListLast(link))
// 	fmt.Println(ListLast(link2))
// }
