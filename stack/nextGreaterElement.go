package stack

//给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集 (既然是子集，说明mums1的元素在nums2都会出现)。
//
//请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
//
//示例 1:
//
//输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出: [-1,3,-1]
//解释:
//对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
//对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
//对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
//
//示例 2:
//输入: nums1 = [2,4], nums2 = [1,2,3,4].
//输出: [3,-1]
//解释:
//对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
//对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出-1 。
//
//提示：
//
//1 <= nums1.length <= nums2.length <= 1000
//0 <= nums1[i], nums2[i] <= 10^4
//nums1和nums2中所有整数 互不相同
//nums1 中的所有整数同样出现在 nums2 中

// 未精简版本
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}
	m := make(map[int]int, len(nums1))
	for k, v := range nums1 {
		m[v] = k
	}

	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		top := stack[len(stack)-1]
		if nums2[i] < nums2[top] {
			stack = append(stack, i)
		} else if nums2[i] == nums2[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && nums2[i] > nums2[top] {
				if v, ok := m[nums2[top]]; ok {
					res[v] = nums2[i]
				}
				stack = stack[:len(stack)-1] // 栈顶出栈
				if len(stack) != 0 {         // 栈还有值，更新top
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}

// 精简版本
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}
	mp := map[int]int{}
	for i, v := range nums1 {
		mp[v] = i
	}
	// 单调栈
	stack := []int{}
	stack = append(stack, 0)

	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {

			top := stack[len(stack)-1]

			if _, ok := mp[nums2[top]]; ok { // 看map里是否存在这个元素
				index := mp[nums2[top]] // 根据map找到nums2[top] 在 nums1中的下标
				res[index] = nums2[i]
			}

			stack = stack[:len(stack)-1] // 出栈
		}
		stack = append(stack, i)
	}
	return res
}
