package dp

import "math"

//给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
//
//示例 1:
//
//输入: amount = 5, coins = [1, 2, 5]
//输出: 4
//解释: 有四种方式可以凑成总金额:
//
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1
//示例 2:
//
//输入: amount = 3, coins = [2]
//输出: 0
//解释: 只用面额2的硬币不能凑成总金额3。
//示例 3:
//
//输入: amount = 10, coins = [10]
//输出: 1
//注意，你可以假设：
//
//0 <= amount (总金额) <= 5000
//1 <= coin (硬币面额) <= 5000
//硬币种类不超过 500 种
//结果符合 32 位符号整数

//在求装满背包有几种方案的时候，认清遍历顺序是非常关键的。
//组合不强调元素之间的顺序，排列强调元素之间的顺序
//如果求组合数就是外层for循环遍历物品，内层for遍历背包。
//如果求排列数就是外层for遍历背包，内层for循环遍历物品。

// 零钱兑换
func change(amount int, coins []int) int {
	// 定义dp数组
	dp := make([]int, amount+1)
	// 初始化,0大小的背包, 当然是不装任何东西了, 就是1种方法
	dp[0] = 1
	// 遍历顺序
	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			// 推导公式
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// 零钱兑换进阶版
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
//
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 示例 2：
//
// 输入：coins = [2], amount = 3
// 输出：-1
// 示例 3：
//
// 输入：coins = [1], amount = 0
// 输出：0
// 示例 4：
//
// 输入：coins = [1], amount = 1
// 输出：1
// 示例 5：
//
// 输入：coins = [1], amount = 2
// 输出：2
// 提示：
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2^31 - 1
// 0 <= amount <= 10^4

//动规五部曲分析如下：
//确定dp数组以及下标的含义
//dp[j]：凑足总额为j所需钱币的最少个数为dp[j]
//确定递推公式
//凑足总额为j - coins[i]的最少个数为dp[j - coins[i]]，那么只需要加上一个钱币coins[i]即dp[j - coins[i]] + 1就是dp[j]（考虑coins[i]）
//所以dp[j] 要取所有 dp[j - coins[i]] + 1 中最小的。
//递推公式：dp[j] = min(dp[j - coins[i]] + 1, dp[j]);
//dp数组如何初始化
//首先凑足总金额为0所需钱币的个数一定是0，那么dp[0] = 0;
//其他下标对应的数值呢？
//考虑到递推公式的特性，dp[j]必须初始化为一个最大的数，否则就会在min(dp[j - coins[i]] + 1, dp[j])比较的过程中被初始值覆盖。
//所以下标非0的元素都是应该是最大值。

//本题求钱币最小个数，那么钱币有顺序和没有顺序都可以，都不影响钱币的最小个数
//本题钱币数量可以无限使用，那么是完全背包。所以遍历的内循环是正序

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 初始化为math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp,j,i)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 版本二,先遍历背包,再遍历物品
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 遍历背包,从1开始
	for j := 1; j <= amount; j++ {
		// 初始化为math.MaxInt32
		dp[j] = math.MaxInt32
		// 遍历物品
		for i := 0; i < len(coins); i++ {
			if j >= coins[i] && dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
