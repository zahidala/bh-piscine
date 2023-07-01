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

func ListAt(l *NodeL, pos int) *NodeL {
	it := l
	count := 0
	for it != nil {
		if count == pos {
			return it
		}
		count = count + 1
		it = it.Next
	}
	return nil
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "hello")
// 	ListPushBack(link, "how are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)

// 	fmt.Println(ListAt(link.Head, 3).Data)
// 	fmt.Println(ListAt(link.Head, 1).Data)
// 	fmt.Println(ListAt(link.Head, 7))
// }
