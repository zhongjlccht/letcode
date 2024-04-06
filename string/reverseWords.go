package string

//给定一个字符串，逐个翻转字符串中的每个单词。
//
//示例 1：
//输入: "the sky is blue"
//输出: "blue is sky the"
//
//示例 2：
//输入: "  hello world!  "
//输出: "world! hello"
//解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//
//示例 3：
//输入: "a good   example"
//输出: "example good a"
//解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
//不要使用辅助空间，空间复杂度要求为O(1)
//#

func reverseWords(s string) string {
	b := []byte(s)

	// 移除前面、中间、后面存在的多余空格
	slow := 0
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for i < len(b) && b[i] != ' ' { // 复制逻辑
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	b = b[0:slow]

	// 翻转整个字符串
	reverse(b)
	// 翻转每个单词
	last := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b[last:i])
			last = i + 1
		}
	}
	return string(b)
}
