package dp

//给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
//
//示例:
//
//nums = [1, 2, 3]
//target = 4
//所有可能的组合为： (1, 1, 1, 1) (1, 1, 2) (1, 2, 1) (1, 3) (2, 1, 1) (2, 2) (3, 1)
//
//请注意，顺序不同的序列被视作不同的组合。
//
//因此输出为 7

// 完全背包问题，物品可以多次放入背包，在意顺序说明是求排列，不是求组合，外层背包，内层物品
func combinationSum4(nums []int, target int) int {
	//定义dp数组
	dp := make([]int, target+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序, 先遍历背包,再循环遍历物品
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] { // j-nums[i]需要大于等于0
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
