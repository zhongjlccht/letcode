package string

import "fmt"

//给定一个字符串 s，它包含小写字母和数字字符，请编写一个函数，将字符串中的字母字符保持不变，而将每个数字字符替换为number。
//
//例如，对于输入字符串 "a1b2c3"，函数应该将其转换为 "anumberbnumbercnumber"。
//
//对于输入字符串 "a5b"，函数应该将其转换为 "anumberb"
//
//输入：一个字符串 s,s 仅包含小写字母和数字字符。
//
//输出：打印一个新的字符串，其中每个数字字符都被替换为了number
//
//样例输入：a1b2c3
//
//样例输出：anumberbnumbercnumber
//
//数据范围：1 <= s.length < 10000。

func replaceNumber(strByte []byte) string {
	// 查看有多少字符
	numCount, oldSize := 0, len(strByte)
	for i := 0; i < len(strByte); i++ {
		if (strByte[i] <= '9') && (strByte[i] >= '0') {
			numCount++
		}
	}
	// 增加长度
	for i := 0; i < numCount; i++ {
		strByte = append(strByte, []byte("     ")...)
	}
	tmpBytes := []byte("number")
	// 双指针从后遍历
	leftP, rightP := oldSize-1, len(strByte)-1
	for leftP < rightP {
		rightShift := 1
		// 如果是数字则加入number
		if (strByte[leftP] <= '9') && (strByte[leftP] >= '0') {
			for i, tmpByte := range tmpBytes {
				strByte[rightP-len(tmpBytes)+i+1] = tmpByte
			}
			rightShift = len(tmpBytes)
		} else {
			strByte[rightP] = strByte[leftP]
		}
		// 更新指针
		rightP -= rightShift
		leftP -= 1
	}
	return string(strByte)
}

func main() {
	var strByte []byte
	fmt.Scanln(&strByte)

	newString := replaceNumber(strByte)

	fmt.Println(newString)
}
