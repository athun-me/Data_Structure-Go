package main

import (
	"fmt"
)

func main() {
	arr := [6]int{5, 8, 2, 9}
	value := 10
	num1, num2 := test(arr, value)
	fmt.Printf("num1 : %d \nnum2 : %d\n", num1, num2)
}

func test(arr [6]int, value int) (int, int) {

	mySet := make(map[int]struct{}) //In go we can use map for hash set

	for i := 0; i < len(arr)-1; i++ {

		num := arr[i]
		match := value - num

		if _, ok := mySet[match]; ok {
			return match, num
		} else {
			mySet[num] = struct{}{}
		}

	}
	return 0, 0
}
