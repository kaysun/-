package main

import (
	"fmt"
	"strings"
)

/*
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/length-of-last-word
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {
	lengthOfLastWord("hello world")
	lengthOfLastWord("hello")
	lengthOfLastWord(" world")
	lengthOfLastWord(" ")
	lengthOfLastWord("")

	lengthOfLastWord2("hello world")
	lengthOfLastWord2("hello")
	lengthOfLastWord2(" world")
	lengthOfLastWord2(" ")
	lengthOfLastWord2("")
}

/*
执行用时 : 0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了25.00%的用户
 */
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	length := len(arr)
	var lenOfLast int
	if length > 0 {
		lastStr := arr[length - 1]
		lenOfLast = len(lastStr)
	} else {
		lenOfLast = 0
	}
	fmt.Println(lenOfLast)
	return lenOfLast
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了75.00%的用户
 */
func lengthOfLastWord2(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	var lenOfLast int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			lenOfLast++
		} else {
			break
		}
	}
	return lenOfLast
}