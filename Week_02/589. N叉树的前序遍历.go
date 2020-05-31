package Week_02

//给定一个 N 叉树，返回其节点值的前序遍历。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其前序遍历: [1,3,5,6,2,4]。
//
//
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var res []int

func preorder(root *Node) []int {
	res = []int{}
	helper(root)
	return res
}

func helper(root *Node) {
	if root != nil {
		res = append(res, root.Val)
		for _, child := range root.Children {
			helper(child)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
