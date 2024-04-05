package dp

import "fmt"

// 完全背包，物品可以放入背包多次

//完全背包
//有N件物品和一个最多能背重量为W的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。
//每件物品都有无限个（也就是可以放入背包多次），求解将哪些物品装入背包里物品价值总和最大。
//
//完全背包和01背包问题唯一不同的地方就是，每种物品有无限件。

// test_CompletePack1 先遍历物品, 在遍历背包
func test_CompletePack1(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ { // j >= weight[i]从循环初始化就天然满足了
			// 推导公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

// test_CompletePack2 先遍历背包, 在遍历物品
func test_CompletePack2(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	// j从0 开始
	for j := 0; j <= bagWeight; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] { // 这里是为了保证 j-weight[i] 大于等于0，不然维度变负了没意义
				// 推导公式
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

func main() {
	weight := []int{1, 3, 4}
	price := []int{15, 20, 30}
	fmt.Println(test_CompletePack1(weight, price, 4))
	fmt.Println(test_CompletePack2(weight, price, 4))
}
