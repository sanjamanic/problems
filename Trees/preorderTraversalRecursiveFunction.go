package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var a []int
	s := make([]int, 1)
	if root != nil {
		s[0] = root.Val
		if root.Left != nil {
			a = preorderTraversal(root.Left)
			for j := 0; j < len(a); j++ {
				s = append(s, a[j])
			}
		}
		if root.Right != nil {
			a = preorderTraversal(root.Right)
			for j := 0; j < len(a); j++ {
				s = append(s, a[j])
			}
		}
		return s
	} else {
		return nil
	}
}
