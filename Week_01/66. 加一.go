package Week_01
//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
// 示例 1:
//
// 输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//
//
// 示例 2:
//
// 输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。
//
// Related Topics 数组


//leetcode submit region begin(Prohibit modification and deletion)
func plusOne(digits []int) []int {
	//1.本题考点在于最高位进位的特殊情形

	//2.准备参数
	pop := 0    //进位
	digits[len(digits)-1] += 1

	for i:= len(digits)-1; i>=0; i--{
		digits[i] += pop
		pop = digits[i]/10
		digits[i] = digits[i]%10

		//没有进位，则没有必要继续遍历
		if pop == 0{
			break
		}
	}

	if digits[0] == 0{
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}
//leetcode submit region end(Prohibit modification and deletion)

