package array

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。
//
//示例 2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
//
//你不需要考虑数组中超出新长度后面的元素。

// 双指针法
func removeElement(nums []int, val int) int {
	length := len(nums)
	res := 0
	for i := 0; i < length; i++ { // i是快指针，遍历元素
		if nums[i] != val {
			nums[res] = nums[i] // res是慢指针，记录覆盖更新的位置，同时也是新数组的长度记录
			res++
		}
	}
	nums = nums[:res]
	return res
}
