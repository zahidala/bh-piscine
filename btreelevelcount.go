package piscine

// package main

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}

// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 		root.Left.Parent = root
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 		root.Right.Parent = root
// 	}
// 	return root
// }

func BTreeLevelCount(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}
	lefth := BTreeLevelCount(root.Left)
	righth := BTreeLevelCount(root.Right)

	if lefth > righth {
		return lefth + 1
	} else {
		return righth + 1
	}
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	fmt.Println(BTreeLevelCount(root))
// }
