package dp

//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
//说明：
//
//拆分时可以重复使用字典中的单词。
//
//你可以假设字典中没有重复的单词。
//
//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false

// 背包法1 dp[i] : 字符串长度为i的话，dp[i]为true，表示可以拆分为一个或多个在字典中出现的单词。
// 如果确定dp[j] 是true，且 [j, i] 这个区间的子串出现在字典里，那么dp[i]一定是true。（j < i ）。
//
// 所以递推公式是 if([j, i] 这个区间的子串出现在字典里 && dp[j]是true) 那么 dp[i] = true。
// 顺序有关系，所以先背包后物品
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 背包法
// 转化为 求装满背包s的前几位字符的方式有几种
func wordBreakV2(s string, wordDict []string) bool {
	// 装满背包s的前几位字符的方式有几种
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i <= len(s); i++ { // 背包
		for j := 0; j < len(wordDict); j++ { // 物品
			if i >= len(wordDict[j]) && wordDict[j] == s[i-len(wordDict[j]):i] {
				dp[i] += dp[i-len(wordDict[j])]
			}
		}
	}

	return dp[len(s)] > 0
}
