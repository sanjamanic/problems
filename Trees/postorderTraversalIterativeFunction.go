package trees

// TODO: IMPROVE!!

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var cnt int
	var rootCurr *TreeNode
	var flag, flag2 bool = false, false
	var s []int
	var sp []*TreeNode
	if root != nil {
		cnt = 0
		rootCurr = root
		for {
			for {
				if rootCurr.Left == nil {
					flag = true
				} else {
					flag = CheckNodeInArray(rootCurr.Left, cnt, sp)
				}
				if flag {
					if rootCurr.Right != nil {
						flag2 = CheckNodeInArray(rootCurr.Right, cnt, sp)
						flag2 = !flag2
						if flag2 {
							rootCurr = rootCurr.Right
							break
						} else {
							cnt, s, sp = AddNodeDecission(rootCurr, cnt, s, sp)
							rootCurr = root
							break
						}
					} else {
						cnt, s, sp = AddNodeDecission(rootCurr, cnt, s, sp)
						rootCurr = root
						break
					}
				} else {
					rootCurr = rootCurr.Left
					break
				}

			}

			if cnt != 0 {
				if root == sp[cnt-1] {
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
