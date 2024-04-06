package backtracking

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 这里一张图，显示的是手机键盘的电话按键，比如2，对应的字母是 abc, 3对应 def , 7对应pqrs

var (
	m    []string
	path []byte
	res  []string
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, res = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return res
	}
	dfs(digits, 0)
	return res
}
func dfs(digits string, start int) {
	if len(path) == len(digits) { //终止条件，字符串长度等于digits的长度
		tmp := string(path)
		res = append(res, tmp) //完成一次搜索
		return
	}
	digit := int(digits[start] - '0') // 将index指向的数字转为int（确定下一个数字）
	str := m[digit-2]                 // 取数字对应的字符集（注意和map中的对应）
	for j := 0; j < len(str); j++ {
		path = append(path, str[j])
		dfs(digits, start+1)
		path = path[:len(path)-1] //将path回退到上一层，这样才会继续推进j
	}
}
