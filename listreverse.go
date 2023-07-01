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

func ListReverse(l *List) {
	var prev *NodeL
	curr := l.Head
	var Next *NodeL
	for curr != nil {

		// Store next
		Next = curr.Next

		// Reverse current node's pointer
		curr.Next = prev

		// Move pointers one position ahead
		prev = curr
		curr = Next
	}
	temp := l.Head
	l.Head = prev
	l.Tail = temp
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, 1)
// 	ListPushBack(link, 2)
// 	ListPushBack(link, 3)
// 	ListPushBack(link, 4)

// 	ListReverse(link)

// 	it := link.Head

// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}

// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)
// }
