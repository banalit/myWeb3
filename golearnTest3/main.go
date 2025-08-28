package golearntest3

import "fmt"

// 有效括号
func main() {
	var test = "()[]{}"
	var test2 = "{(})"
	valid(test)
	valid(test2)
}

func valid(test string) {
	fmt.Println("test is ", test)

}
