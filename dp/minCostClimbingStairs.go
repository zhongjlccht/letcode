package dp

//给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
//你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
//请你计算并返回达到楼梯顶部的最低花费。

//示例 1：
//
//输入：cost = [10, 15, 20]
//输出：15
//解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。
//示例 2：
//
//输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
//输出：6
//解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。
//提示：
//
//cost 的长度范围是 [2, 1000]。
//cost[i] 将会是一个整型数据，范围为 [0, 999] 。

// 最小代价爬楼梯
func minCostClimbingStairs(cost []int) int {
	f := make([]int, len(cost)+1)
	f[0], f[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		f[i] = min(f[i-1]+cost[i-1], f[i-2]+cost[i-2])
	}
	return f[len(cost)]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
