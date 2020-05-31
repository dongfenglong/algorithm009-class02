package Week_02

//给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其层序遍历:
//
// [
//     [1],
//     [3,2,4],
//     [5,6]
//]
//
//
//
//
// 说明:
//
//
// 树的深度不会超过 1000。
// 树的节点总数不会超过 5000。
// Related Topics 树 广度优先搜索


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var res [][]int
var m map[int][]int
func levelOrder(root *Node) [][]int {

	m = make(map[int][]int)
	helper3(root, 0)
	res = make([][]int, len(m))
	for l, nums := range m {
		res[l] = nums
	}
	return res
}

func helper3(root *Node, level int){
	if root!=nil{
		m[level] = append(m[level], root.Val)

		for _, child := range root.Children{
			helper(child, level+1)
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)

