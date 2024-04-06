package stack

//给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，
//这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
//
//示例 1:
//
//输入: [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；数字 2 找不到下一个更大的数；第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//提示:
//
//1 <= nums.length <= 10^4
//-10^9 <= nums[i] <= 10^9

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	for i := 0; i < len(result); i++ {
		result[i] = -1
	}
	//单调递减，存储数组下标索引
	stack := make([]int, 0)
	for i := 0; i < length*2; i++ {
		for len(stack) > 0 && nums[i%length] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			result[index] = nums[i%length]
		}
		stack = append(stack, i%length)
	}
	return result
}
