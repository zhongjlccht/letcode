package string

//实现 strStr() 函数。
//
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//
//示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2
//
//示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1
//
//说明: 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 对于本题而言，当 needle 是空字符串时我们应当返回 0 。
//这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
//
//#

//考 KMP 和 前缀表
// 前缀表：记录下标i之前（包括i）的字符串中，有多大长度的相同前缀后缀
//字符串的前缀是指不包含最后一个字符的所有以第一个字符开头的连续子串
//后缀是指不包含第一个字符的所有以最后一个字符结尾的连续子串

// 方法一:前缀表使用减1实现

// getNext 构造前缀表next
// params:
//
//	next 前缀表数组
//	s 模式串
func getNextV1(next []int, s string) {
	j := -1 // j表示 最长相等前后缀长度
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j] // 回退前一位
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
}
func strStrV1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	getNextV1(next, needle)
	j := -1 // 模式串的起始位置 next为-1 因此也为-1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // 寻找下一个匹配点
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 { // j指向了模式串的末尾
			return i - len(needle) + 1
		}
	}
	return -1
}

// 方法二: 前缀表无减一或者右移

// getNext 构造前缀表next
// params:
//
//	next 前缀表数组
//	s 模式串
func getNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}
