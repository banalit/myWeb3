package task2

import "fmt"

func PtrAdd10(num *int) {
	fmt.Println("num :", *num)
	*num = *num + 10
}

func sliceMulti2(num *[]int) {
	for i := 0; i < len(*num); i++ {
		(*num)[i] = (*num)[i] * 2
	}
}
