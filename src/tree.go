package src

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func CheckBST(node *Node) bool {
	return checkTree(node, -10000, 10000)
}

func checkTree(node *Node, min int, max int) bool {
	if node == nil {
		return true
	}
	return min < node.Data && node.Data < max && checkTree(node.Left, min, node.Data) && checkTree(node.Right, node.Data, max)
}