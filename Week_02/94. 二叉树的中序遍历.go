package Week_02

//给定一个二叉树，返回它的中序 遍历。
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
//输出: [1,3,2]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表

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

func inorderTraversal(root *TreeNode) []int {
	res = []int{}

	_helper(root)
	return res
}

func _helper(root *TreeNode) {
	if root != nil {
		helper(root.Left)
		res = append(res, root.Val)
		helper(root.Right)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
