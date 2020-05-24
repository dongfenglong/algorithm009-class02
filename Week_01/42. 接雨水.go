package Week_01
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Mar
//cos 贡献此图。
//
// 示例:
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
// Related Topics 栈 数组 双指针


//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	if len(height)<=2{
		return 0
	}

	left, right, max := 0, len(height)-1, 0
	area := 0
	for id, val := range height{
		if val > height[max]{
			max=id
		}
	}

	for i:=0;i< max;i++{
		if height[i] > height[left]{
			left = i
		}else{
			area += height[left]-height[i]
		}
	}

	for i:= len(height)-1;i>max;i--{
		if height[i] > height[right]{
			right = i
		}else{
			area += height[right]-height[i]
		}
	}
	return area
}
//leetcode submit region end(Prohibit modification and deletion)

