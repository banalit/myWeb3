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
	// bytes := []byte(str)
	bytes := str
	// fmt.Println("test is ", test)
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
		fmt.Println(test, " not 回文数")
	}
	return check
}

func main() {
	isPalindrome(123321)
	isPalindrome(123123)
	isPalindrome(123123999)
	isPalindrome(-123321)
}
