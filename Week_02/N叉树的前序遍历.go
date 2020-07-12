/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
var ans []int

func preorder(root *Node) []int {
	ans = []int{}
	if root == nil {
		return ans
	}
	helper(root)
	return ans
}

func helper(root *Node) {
	ans = append(ans, root.Val)
	if root.Children != nil {
		for _, child := range root.Children {
			helper(child)
		}
	}
}