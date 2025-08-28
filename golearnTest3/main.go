package golearntest3

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

}
