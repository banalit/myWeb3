package main

import (
	"fmt"
	"strconv"
)

// 回文数
func doCheck(test int) bool {
	if test < 0 {
		return false
	}
	if test%10 == 0 {
		return false
	}
	if test < 10 {
		return true
	}
	str := strconv.Itoa(test)
	bytes := []byte(str)
	fmt.Println("test is ", test)
	for i := 0; i < len(bytes)/2; i++ {
		b := bytes[i]
		e := bytes[len(bytes)-1-i]
		if b != e {
			return false
		}
	}
	return true
}

func isPalindrome(test int) bool {
	check := doCheck(test)
	if check {
		fmt.Println(test, " is 回文数")
	} else {
		fmt.Println(test, " is not 回文数")
	}
	return check
}

func main() {
	x := 123321
	y := 123123
	isPalindrome(x)
	isPalindrome(y)
}
