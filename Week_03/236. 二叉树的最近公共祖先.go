package Week_03

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//LeetCode官方题解
	//1.通过遍历，建立 子节点 -> 父结点 的hash映射
	//2.通过建立的hash映射，让p结点向上访问，寻找祖先结点，并将这些祖先节点记录到set中（set中包含p结点）
	//3.让q结点向上访问祖先结点，如果有祖先结点在set中出现，就返回该结点


	fatherMap := make(map[*TreeNode]*TreeNode)
	fatherSet := make(map[*TreeNode]*TreeNode)
	//1.
	fatherMap[root] = nil
	_search(root, fatherMap)

	//2.
	for p!=nil{
		fatherSet[p] = nil
		p = fatherMap[p]
	}

	//3.
	for q!=nil{
		if _, ok := fatherSet[q]; ok{
			return q
		}
		q = fatherMap[q]
	}

	return nil
}

func _search(root *TreeNode, fatherMap map[*TreeNode]*TreeNode){
	if root == nil{
		return
	}

	if root.Left != nil{
		fatherMap[root.Left] = root
		_search(root.Left, fatherMap)
	}

	if root.Right != nil{
		fatherMap[root.Right] = root
		_search(root.Right, fatherMap)
	}
}
