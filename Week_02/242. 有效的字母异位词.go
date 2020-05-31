package Week_02

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
//输出: false
//
// 说明:
//你可以假设字符串只包含小写字母。
//
// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 排序 哈希表


//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	//两个字符串长度相同，如果不是字母异位词，那么在第二个字符串遍历递减时，肯定会出现字符为负值的情况

	keys := [26]int{}
	for _, key := range s {
		keys[key-'a']++
	}

	for _, key := range t{
		if keys[key-'a']> 0{
			keys[key-'a']--
		}else{
			return false
		}
	}

	return true
}
//leetcode submit region end(Prohibit modification and deletion)