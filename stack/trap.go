package stack

// 接雨水

// 暴力法，左右双向朝中间遍历 按列加总
func trap1(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right { // 2者不相等，说明还有列没有遍历完
		if height[left] < height[right] { // 如果左侧柱子低，以左侧柱子高度来减，进入条件分支处理
			if height[left] >= leftMax { //没有出现凹槽，则更新左侧最大柱子高度
				leftMax = height[left]
			} else { // 出现了凹槽，计算此列雨水高度
				res += leftMax - height[left] // //右边必定有柱子挡水，所以遇到所有值小于等于leftMax的，全部加入水池中
			}
			left++
		} else { // 如果右侧柱子低，以右侧柱子高度来减，进入条件分支处理，相等的话，则2者都可以
			if height[right] > rightMax {
				rightMax = height[right] // //设置右边最高柱子
			} else {
				res += rightMax - height[right] // //左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
			}
			right--
		}
	}
	return res
}

// 双指针， 按列加总
func trap2(height []int) int {
	sum := 0
	n := len(height)
	lh := make([]int, n)
	rh := make([]int, n)
	lh[0] = height[0]
	rh[n-1] = height[n-1]
	for i := 1; i < n; i++ { // 从下标1开始，从i位置往左边看的左侧柱子中的最大高度
		lh[i] = max(lh[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- { // 从下标n-2, 即倒数第2个位置，往右边看的右侧柱子中的最大高度
		rh[i] = max(rh[i+1], height[i])
	}
	for i := 1; i < n-1; i++ { // 遍历每个位置，如果是凹槽位置，有最大左侧和最大右侧柱子，看2者哪个最小，再跟当前i的高度相减去，就是当前列接的雨水
		h := min(rh[i], lh[i]) - height[i] // 如果左侧没有柱子挡，或者右侧没有柱子挡，那变成0-height[i]，结果是<=0，
		if h > 0 {                         //只要>0，就说明是取的是凹槽位置的雨水
			sum += h
		}
	}
	return sum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 单调栈法， 按行加总
func trap3(height []int) int {
	if len(height) <= 2 { // 只有2列，形成不了凹槽
		return 0
	}
	st := make([]int, 1, len(height)) // 切片模拟单调栈，st存储的是高度数组下标
	var res int
	for i := 1; i < len(height); i++ {
		if height[i] < height[st[len(st)-1]] { //st[len(st)-1]是栈顶位置的值，这个值是下标
			st = append(st, i) // 只要小于栈顶，就将下标入栈，形成凹槽的左侧
		} else if height[i] == height[st[len(st)-1]] {
			st = st[:len(st)-1] // 比较的新元素和栈顶的元素相等，去掉栈中的，入栈新元素下标；
			st = append(st, i)  //这里是因为当出现右侧柱子时，（i - st[len(st)-1] - 1）会取左侧比当前高度高那一列的下标计算宽度，会将出栈列的宽度也算进去，高度一样的列去掉不影响计算面积
		} else {
			// for 循环会把凹槽，当前右侧柱子，比左侧高的全部计算面积
			for len(st) != 0 && height[i] > height[st[len(st)-1]] { // height[i]作为右侧柱子，比栈顶下标代表柱子的高度高，说明出现凹槽
				top := st[len(st)-1] //读取栈顶存的下标值
				st = st[:len(st)-1]  //将栈顶元素出栈
				if len(st) != 0 {    //说明左侧有柱子挡水
					// min(height[st[len(st)-1], height[i]]) 是比较左右2侧柱子，取高度小的，然后减去 height[top]，得到的是长度，然后面积是长度*宽度
					tmp := (min(height[st[len(st)-1]], height[i]) - height[top]) * (i - st[len(st)-1] - 1) //  (i - st[len(st)-1] - 1)这个计算的是宽度
					res += tmp
				}
			}
			// 当前的i都要放进去，因为凹槽右侧可能有更高的右侧柱子比当前height[i]高
			st = append(st, i)
		}
	}
	return res
}

// 单调栈压缩版   ---- 这个版本最好理解，按行累加，就背这个-----
func trap4(height []int) int {
	stack := make([]int, 0)
	res := 0

	// 无需事先将第一个柱子的坐标入栈，因为它会在该for循环的最后入栈
	for i := 0; i < len(height); i++ {
		// 满足栈不为空并且当前柱子高度大于栈顶对应的柱子高度的情况时
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			// 获得凹槽高度
			mid := height[stack[len(stack)-1]]
			// 凹槽坐标出栈
			stack = stack[:len(stack)-1]

			// 如果栈不为空则此时栈顶元素为左侧柱子坐标
			if len(stack) > 0 {
				// 求得雨水高度
				h := min(height[i], height[stack[len(stack)-1]]) - mid
				// 求得雨水宽度
				w := i - stack[len(stack)-1] - 1
				res += h * w
			}
		}
		// 如果栈为空或者当前柱子高度小于等于栈顶对应的柱子高度时入栈
		stack = append(stack, i)
	}
	return res
}
