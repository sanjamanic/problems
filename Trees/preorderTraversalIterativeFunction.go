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
func preorderTraversal(root *TreeNode) []int {
	var cnt, cnth int
	var rootCurr *TreeNode
	var flag bool = false
	s := make([]int, 1) // solution
	sp := make([]*TreeNode, 1)
	ah := make([]int, 1) // help array
	ahp := make([]*TreeNode, 1)
	if root != nil {
		s[0] = root.Val
		sp[0] = root
		cnt = 1
		cnth = 0
		rootCurr = root
		for {
			for {
				if rootCurr.Left == nil {
					break
				}
				cnt, s, sp = AddNodeDecission(rootCurr.Left, cnt, s, sp)
				flag = CheckNodeInArray(rootCurr.Left, cnth, ahp)
				if flag {
					break
				}
				rootCurr = rootCurr.Left
			}

			if rootCurr.Right != nil {
				cnt, s, sp = AddNodeDecission(rootCurr.Right, cnt, s, sp)
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
			if root == ahp[cnth] {
				break
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
