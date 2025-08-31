package main

import (
	"fmt"
)

//大整数加1

func main() {
	digitals := []uint{1, 2, 3, 9, 0}
	addAndPrint(digitals)
	digitals = []uint{0, 2, 3, 9, 9}
	addAndPrint(digitals)
}

func addAndPrint(digs []uint) {
	fmt.Print(digs, ":")
	fmt.Println(addOne(digs))
}
func addOne(digs []uint) []uint {
	llen := len(digs)
	for i := llen - 1; i >= 0; i-- {
		if digs[i] != 9 {
			digs[i]++
			for j := i + 1; j < llen; j++ {
				digs[j] = 0
			}
			return digs
		}
	}
	t := make([]uint, llen+1)
	t[0] = 1
	return t
}
