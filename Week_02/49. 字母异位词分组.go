package Week_02

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}

	//1.定义质数数组: primes
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	//2.定义数组map: strMap
	strMap := make(map[int][]string)
	//3.定义异位词分组数组: strArr
	strArr := make([][]string, 0)

	//4.遍历字符串数组
	for _, str := range strs {
		//4.1遍历字符串，与质数数组对应下标的值 相乘
		num := 1
		for _, s := range str {
			num *= primes[s-'a']
		}

		strMap[num] = append(strMap[num], str)
	}

	for _, arr := range strMap {
		strArr = append(strArr, arr)
	}

	return strArr
}

//leetcode submit region end(Prohibit modification and deletion)
