/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	candidates := []*TreeNode{root}
	for i := 0; i < len(candidates); i++ { 
		cand := candidates[i]
		if cand.Val != subRoot.Val {
			if cand.Left != nil {
				candidates = append(candidates, cand.Left)
			}
			if cand.Right != nil {
				candidates = append(candidates, cand.Right)
			}	
		}else {
			if subtree(cand, subRoot) {
				return true
			}
			if cand.Left != nil {
				candidates = append(candidates, cand.Left)
			}
			if cand.Right != nil {
				candidates = append(candidates, cand.Right)
			}	
		}
	}
	return false
}

func subtree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val{
		return subtree(a.Left, b.Left) && subtree(a.Right, b.Right)
	}
	return false
}
