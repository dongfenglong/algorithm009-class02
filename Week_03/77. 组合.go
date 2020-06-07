package Week_03

var res [][]int

func combine(n int, k int) [][]int {
	if n < 1 || n < k {
		return nil
	}
	//1.初始化返回值
	res = [][]int{}

	//2.递归
	_genarate(n, k, 1, []int{})
	return res

}

//nums 用于存放选中的数字
func _genarate(n, k, num int, nums []int) {
	//1.逻辑终结条件
	if num > n {
		return
	}
	if len(nums) == k {
		res = append(res, nums)
		return
	}

	//2.处理本层逻辑
	if num <= n {
		//当前数字不添加到数组中
		_genarate(n, k, num+1, nums)
		//当前数字添加到数组中
		_genarate(n, k, num+1, append(nums, num))
	}

	//3.下探到下一层
	//4.清理本层状态
}
