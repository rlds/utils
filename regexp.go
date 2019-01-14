package utils

import (
 	"regexp"
)

// 是否匹配正则
func IsMatch(strText, strReg string) bool {
	reg := regexp.MustCompile(strReg)
	return reg.Match([]byte(strText))
}

// 正则表达获取字符串
func MatchData(strText, strReg string) string {
	reg := regexp.MustCompile(strReg)
	arrMatch := reg.FindAllStringSubmatch(strText, -1)
	result := ""
	if len(arrMatch) > 0 {
		result = arrMatch[0][1]
	}
	return result
}

// 正则返回多行
func MatchMutilData(strText, strReg string) [][]string {
	reg := regexp.MustCompile(strReg)
	return reg.FindAllStringSubmatch(strText, -1)
}

// 正则表达式提取重复单列数据
func MatchSingleLine(strText, strReg string) []string {
	rows := []string{}
	reg := regexp.MustCompile(strReg)
	arrMatch := reg.FindAllStringSubmatch(strText, -1)
	if len(arrMatch) > 0 {
		for _, val := range arrMatch {
			for i:=1; i<len(val); i++ {
				rows = append(rows, val[i])
			}
		}
	}
	return rows
}

// 正则表达式提取重复多列数据
// 从列匹配转换成行
func MatchMutilLine(strText, strReg string) [][]string {
	reg := regexp.MustCompile(strReg)
	arrMatch := reg.FindAllStringSubmatch(strText, -1)
	rows := make([][]string, len(arrMatch))
	if len(arrMatch) > 0 {
		for i := 0; i < len(arrMatch); i++ {
			rows[i] = make([]string, 0)
			for j := 1; j < len(arrMatch[i]); j++ {
				rows[i] = append(rows[i], arrMatch[i][j])
			}
		}
	}
	return rows
}



