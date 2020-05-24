package Week_01
//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy

	for l1!=nil&&l2!=nil{
		if l1.Val <= l2.Val{
			pre.Next, pre, l1 = l1, l1, l1.Next
		}else{
			pre.Next, pre, l2 = l2, l2, l2.Next
		}
	}

	if l1 == nil{
		pre.Next = l2
	}else{
		pre.Next = l1
	}
	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)

