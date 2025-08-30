package main

import "fmt"

// 最长公共前缀
func main() {
	test := []string{"abc", "abefe", "ab11232"}
	fmt.Println(test, ":", getLongestPrefix(test))
	test = []string{"你好", "你好世界", "你好大家"}
	// r []rune:=[]rune('1')  //会报错？？
	// fmt.Println(r)
	fmt.Println(test, ":", getLongestPrefix(test))
}

func getMinLen(test []string) int {
	minLen := len([]rune(test[0]))
	for i := 0; i < len(test); i++ {
		if minLen > len([]rune(test[i])) {
			minLen = len([]rune(test[i]))
		}
	}
	return minLen
}

func getLongestPrefix(test []string) string {
	var minLen = getMinLen(test)
	var runes []rune
	for i := 0; i < minLen; i++ {
		r := []rune(test[0])[i]
		for j := 0; j < len(test); j++ {
			testRune := []rune(test[j])
			if r != testRune[i] {
				if i == 0 {
					return ""
				}
				return string(runes)
			}
		}
		runes = append(runes, r)
	}
	return string(runes)
}
