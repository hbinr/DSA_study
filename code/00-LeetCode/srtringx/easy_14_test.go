package stringx

import (
	"fmt"
	"strings"
	"testing"
)

/*
	https://leetcode-cn.com/problems/longest-common-prefix/

	编写一个函数来查找字符串数组中的最长公共前缀。

	如果不存在公共前缀，返回空字符串 ""。



	示例 1：
	输入：strs = ["flower","flow","flight"]
	输出："fl"

	分析：
	https://www.geekxh.com/1.0.%E6%95%B0%E7%BB%84%E7%B3%BB%E5%88%97/002.html#_02%E3%80%81%E9%A2%98%E8%A7%A3%E5%88%86%E6%9E%90

*/

func TestLCP(t *testing.T) {
	fmt.Println("res： ", longestCommonPrefix([]string{"flower", "flow", "flight"}))

}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs {
		// Index返回s中prefix的第一个实例的索引；如果s中不存在prefix，则返回-1。
		for strings.Index(s, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			// 截取掉基准元素prefix的最后一个元素，再次和s进行比较，直至满足 strings.Index(s, prefix) == 0
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
