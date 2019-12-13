package trees

// TODO: IMPROVE

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var cnt, cnth int
	var rootCurr *TreeNode
	var flag bool = false
	var s, ah []int
	var sp, ahp []*TreeNode
	if root != nil {
		cnt = 0
		cnth = 0
		rootCurr = root
		for {
			for {
				if rootCurr.Left == nil {
					flag = true
				} else {
					flag = CheckNodeInArray(rootCurr.Left, cnth, ahp)
				}
				if flag {
					cnt, s, sp = AddNodeDecission(rootCurr, cnt, s, sp)
					break
				} else {
					rootCurr = rootCurr.Left
				}

			}

			if rootCurr.Right != nil {
				flag = CheckNodeInArray(rootCurr.Right, cnth, ahp)
				if flag {
					cnth, ah, ahp = AddNodeDecission(rootCurr, cnth, ah, ahp)
					rootCurr = root
				} else {
					rootCurr = rootCurr.Right
				}
			} else {
				cnth, ah, ahp = AddNodeDecission(rootCurr, cnth, ah, ahp)
				rootCurr = root
			}
			if cnth != 0 {
				if root == ahp[cnth-1] {
					break
				}
			}
		}
		return s
	} else {
		return nil
	}
}

func AddNodeDecission(root *TreeNode, cnt int, s []int, sp []*TreeNode) (int, []int, []*TreeNode) {
	var flag bool = true
	if root != nil {
		flag = CheckNodeInArray(root, cnt, sp)
		flag = !flag
	} else {
		flag = false
	}
	if flag {
		cnt++
		s = append(s, root.Val)
		sp = append(sp, root)
	}
	return cnt, s, sp
}

func CheckNodeInArray(root *TreeNode, cnt int, sp []*TreeNode) bool {
	var flag bool = false
	if root != nil {
		for i := 0; i < cnt; i++ {
			if sp[i] == root {
				flag = true
				break
			}
		}
	}
	return flag
}
