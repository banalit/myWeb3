package main

import (
	"fmt"
	"strings"
)

// 有效括号
func main() {
	test := "()[]{}"
	fmt.Println(test, "is valide ", isValid(test))
	test = "{(})"
	fmt.Println(test, "is valide ", isValid(test))

	test = "()[({})]{}"
	fmt.Println(test, "is valide ", isValid(test))
}

func isValid(test string) bool {
	// fmt.Println("test is ", test)
	patterns := []string{"()", "[]", "{}"}
	for i := 0; i < len(test)/2; i++ {
		for _, v := range patterns {
			if len(test) == 0 {
				return true
			}
			if len(test) == 1 {
				return false
			}
			test = strings.ReplaceAll(test, v, "")
		}
	}
	// fmt.Println("last test:", test)
	if len(test) == 0 {
		return true
	}
	return false
}
