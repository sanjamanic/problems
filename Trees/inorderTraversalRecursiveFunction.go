package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var a, s []int
	if root != nil {
		if root.Left != nil {
			a = inorderTraversal(root.Left)
			for j := 0; j < len(a); j++ {
				s = append(s, a[j])
			}
		}
		s = append(s, root.Val)
		if root.Right != nil {
			a = inorderTraversal(root.Right)
			for j := 0; j < len(a); j++ {
				s = append(s, a[j])
			}
		}
		return s
	} else {
		return nil
	}
}
