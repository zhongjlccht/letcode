package stack

//请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
//
//提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

// 单调栈法(未精简版本)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 初始化栈顶元素为第一个下标索引0
	stack := []int{0}

	for i := 1; i < len(temperatures); i++ {
		top := stack[len(stack)-1]
		if temperatures[i] < temperatures[top] {
			stack = append(stack, i)
		} else if temperatures[i] == temperatures[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[top] {
				res[top] = i - top
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}

// 单调递减栈
func dailyTemperatures2(num []int) []int {
	ans := make([]int, len(num))
	stack := []int{}
	for i, v := range num {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}
