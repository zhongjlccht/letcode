package dp

// 01 背包问题
// 有n件物品和一个最多能背重量为w 的背包。
// 第i件物品的重量是weight[i]，得到的价值是value[i] 。
// 每件物品只能用一次，
// 求解将哪些物品装入背包里物品价值总和最大。

// dp[i][j] 表示从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少
// 推导公式：dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i])

// 二维DP解法，dp[i][j] 表示从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少
func test_2_wei_bag_problem1(weight, value []int, bagweight int) int {
	// 定义dp数组，
	// 这里因为go， 默认每个dp[i][j]都被初始化成0了，不然还需要对dp[i][0]初始化成0，因为背包容量0那一列放不了物品，价值都是0
	// 且dp[0][j], j背包容量如果小于物品0的重量，也是放不进去的，也是价值0
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagweight+1) //bagweight+1，因为有背包大小0的一列，比如bagweight传入4，则0/1/2/3/4共5种背包大小的情况
	}
	// 初始化， j >= weight[0]是因为背包容量需要能放进物品0的重量
	for j := bagweight; j >= weight[0]; j-- { // weight[0]是第一件物品重量，这里是推导dp[0][j]的值，因为dp[i][j]依赖dp[0][j]
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// 递推公式
	for i := 1; i < len(weight); i++ { // 先遍历物品重量 i=0的情况，上面初始化已经处理过了，所以i从1开始
		//正序,也可以倒序
		for j := 0; j <= bagweight; j++ { // 然后遍历背包容量，这里j是具体的背包容量值
			if j < weight[i] {
				dp[i][j] = dp[i-1][j] // 物品放不下，价值没有增加
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i]) // 能放下物品i，说明之前的背包容量是[j-weight[i]]，在这基础上增加value[i]
			}
		}
	}
	return dp[len(weight)-1][bagweight]
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test_2_wei_bag_problem1(weight, value, 4) // 二维解法
	test_1_wei_bag_problem(weight, value, 4)  // 一维解法
}

// 第二种解法，将维度，用一维数组来解
// 在一维dp数组中，dp[j]表示：容量为j的背包，所背的物品价值可以最大为dp[j]。
// dp[j]可以通过dp[j - weight[i]]推导出来，dp[j - weight[i]]表示容量为j - weight[i]的背包所背的最大价值。
// dp[j - weight[i]] + value[i] 表示 容量为 j - 物品i重量 的背包 加上 物品i的价值。（也就是容量为j的背包，放入物品i了之后的价值即：dp[j]）
// 此时dp[j]有两个选择，一个是取自己dp[j] 相当于 二维dp数组中的dp[i-1][j]，即不放物品i，一个是取dp[j - weight[i]] + value[i]，即放物品i，指定是取最大的，毕竟是求最大价值，
// 递推公式：dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
// 假设物品价值都是大于0的，所以dp数组初始化的时候，都初始为0就可以了
// 注意！！！：遍历顺序需要倒序遍历背包大小，
func test_1_wei_bag_problem(weight, value []int, bagWeight int) int {
	// 定义 and 初始化
	dp := make([]int, bagWeight+1)
	// 递推顺序
	for i := 0; i < len(weight); i++ { // 遍历物品
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j := bagWeight; j >= weight[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	//fmt.Println(dp)
	return dp[bagWeight]
}
