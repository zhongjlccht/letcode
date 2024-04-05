package array

//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
//
//示例：
//
//输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//提示：
//
//1 <= target <= 10^9
//1 <= nums.length <= 10^5
//1 <= nums[i] <= 10^5

// 双指针， 滑动窗口解法
func minSubArrayLen(target int, nums []int) int {
	i := 0                   // 窗口的左侧下标
	l := len(nums)           // 数组长度
	sum := 0                 // 子数组之和
	result := l + 1          // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ { // 移动窗口的右侧下标
		sum += nums[j]
		for sum >= target { //说明需要缩小窗口
			subLength := j - i + 1 //计算子数组长度
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++ //向右挪动窗口左侧
		}
	}
	if result == l+1 {
		return 0 // 不存在这样的子数组
	} else {
		return result
	}
}
