package hashmap

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
//
//示例 2: 输入: s = "rat", t = "car" 输出: false
//
//说明: 你可以假设字符串只包含小写字母。

// 只遍历一次
func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	records := [26]int{}
	for index := 0; index < len(s); index++ {
		if s[index] == t[index] {
			continue
		}
		sCharIndex := s[index] - 'a'
		records[sCharIndex]++
		tCharIndex := t[index] - 'a'
		records[tCharIndex]--
	}
	for _, record := range records {
		if record != 0 {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		record[r-rune('a')]++
	}
	for _, r := range t {
		record[r-rune('a')]--
	}

	return record == [26]int{} // record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符。
}
