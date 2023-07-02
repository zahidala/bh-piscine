// package piscine

package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root

	// newNode := &TreeNode{Data: data}

	// if newNode.Data < root.Data {
	// 	if root.Left == nil {
	// 		root.Left = newNode
	// 	} else {
	// 		BTreeInsertData(root.Left, newNode.Data)
	// 		root.Left.Parent = root
	// 	}
	// }

	// if newNode.Data > root.Data {
	// 	if root.Right == nil {
	// 		root.Right = newNode
	// 	} else {
	// 		BTreeInsertData(root.Right, newNode.Data)
	// 		root.Left.Parent = root
	// 	}
	// }
	// return newNode
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	fmt.Println(root.Left.Data)
	fmt.Println(root.Data)
	fmt.Println(root.Right.Left.Data)
	fmt.Println(root.Right.Data)

}
