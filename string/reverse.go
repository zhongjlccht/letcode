package string

import "fmt"

//字符串的右旋转操作是把字符串尾部的若干个字符转移到字符串的前面。给定一个字符串 s 和一个正整数 k，请编写一个函数，将字符串中的后面 k 个字符移到字符串的前面，实现字符串的右旋转操作。
//
//例如，对于输入字符串 "abcdefg" 和整数 2，函数应该将其转换为 "fgabcde"。
//
//输入：输入共包含两行，第一行为一个正整数 k，代表右旋转的位数。第二行为字符串 s，代表需要旋转的字符串。
//
//输出：输出共一行，为进行了右旋转操作后的字符串。
//样例输入：
//2
//abcdefg
//样例输出：
//fgabcde
//数据范围：1 <= k < 10000, 1 <= s.length < 10000;

//不能申请额外空间，只能在本串上操作

func reverse33(strByte []byte, l, r int) {
	for l < r {
		strByte[l], strByte[r] = strByte[r], strByte[l]
		l++
		r--
	}
}

func main() {
	var str string
	var target int

	fmt.Scanln(&str)
	fmt.Scanln(&target)
	strByte := []byte(str)

	reverse33(strByte, 0, len(strByte)-1)      // 先整个字符串倒序
	reverse33(strByte, 0, target-1)            // 前面子串倒序
	reverse33(strByte, target, len(strByte)-1) //后面子串倒序 最终负负得正

	fmt.Printf(string(strByte))
}
