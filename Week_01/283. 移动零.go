package Week_01
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针


//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int)  {
	//0.判断nums长度
	if len(nums) <= 1{
		return
	}

	//1.定义快慢指针
	fast, slow := 0, -1

	for fast < len(nums){
		//1.1 第一个0
		if nums[fast]==0{
			slow=fast
			break
		}
		fast++
	}

	for fast < len(nums){
		//1.2 遇到非零值， 与第一个0交换
		if slow >= 0 && nums[fast] != 0{
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
}
//leetcode submit region end(Prohibit modification and deletion)

