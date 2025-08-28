package main

import "fmt"

// 有效括号
func main() {
	var test = "()[]{}"
	var test2 = "{(})"
	isValid(test)
	isValid(test2)
}

func isValid(test string) {
	fmt.Println("test is ", test)
	// llen :=len(test)
	var code []string = []string{}
	for i, v := range test {
		code[i] = v
		if len(code) > 1 {

		}
	}
}
