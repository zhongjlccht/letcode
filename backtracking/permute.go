package backtracking

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出: [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]
//#

var (
	res1  [][]int
	path1 []int
	st    []bool // state的缩写
)

func permute(nums1 []int) [][]int {
	res1, path1 = make([][]int, 0), make([]int, 0, len(nums1))
	st = make([]bool, len(nums1))
	dfs1(nums1, 0)
	return res1
}

func dfs1(nums1 []int, cur int) {
	if cur == len(nums1) { // 终止条件
		tmp := make([]int, len(path1))
		copy(tmp, path1)
		res1 = append(res1, tmp)
	}
	for i := 0; i < len(nums1); i++ {
		if !st[i] { // 还没出现过
			path1 = append(path1, nums1[i])
			st[i] = true
			dfs1(nums1, cur+1)
			st[i] = false
			path1 = path1[:len(path1)-1]
		}
	}
}
