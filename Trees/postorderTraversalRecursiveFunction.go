package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var a, s []int
	if root != nil {
		if root.Left != nil {
			a = postorderTraversal(root.Left)
			for j := 0; j < len(a); j++ {
				s = append(s, a[j])
			}
		}
		if root.Right != nil {
			a = postorderTraversal(root.Right)
			for j := 0; j < len(a); j++ {
				s = append(s, a[j])
			}
		}
		s = append(s, root.Val)
		return s
	} else {
		return nil
	}
}
