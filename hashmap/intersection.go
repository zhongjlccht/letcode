package hashmap

//题意：给定两个数组，编写一个函数来计算它们的交集。
//示例1:
//	输入 nums1=[1,2,2,1] nums2=[2,2] 输出[2]
//示例2:
//	输入 nums1=[4,9,5] nums2=[9,4,9,8,4] 输出[9,4]

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0) // 用map模拟set
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
