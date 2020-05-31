package Week_02

//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,2,3]
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []int
func preorderTraversal(root *TreeNode) []int {
	res = []int{}
	helper2(root)
	return res
}

func helper2(root *TreeNode){
	if root!=nil{
		res = append(res, root.Val)
		helper(root.Left)
		helper(root.Right)
	}
}
//leetcode submit region end(Prohibit modification and deletion)

