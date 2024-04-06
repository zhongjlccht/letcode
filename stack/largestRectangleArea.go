package stack

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。

// 输入 heights=[2,1,5,6,2,3] 输出10， 这里5和6 挨在一起的面积最大，5+5=10

func largestRectangleArea(heights []int) int {
	max := 0
	// 使用切片实现栈
	stack := make([]int, 0)
	// 数组头部加入0
	heights = append([]int{0}, heights...)
	// 数组尾部加入0
	heights = append(heights, 0)
	// 初始化栈，序号从0开始
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		// 结束循环条件为：当即将入栈元素<top元素，也就是形成非单调递增的趋势
		for heights[stack[len(stack)-1]] > heights[i] {
			// mid 是top
			mid := stack[len(stack)-1]
			// 出栈
			stack = stack[0 : len(stack)-1]
			// left是top的下一位元素，i是将要入栈的元素
			left := stack[len(stack)-1]
			// 高度x宽度
			tmp := heights[mid] * (i - left - 1)
			if tmp > max {
				max = tmp
			}
		}
		stack = append(stack, i)
	}
	return max
}
