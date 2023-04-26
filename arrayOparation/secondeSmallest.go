package main

import (
	"fmt"
)

func main() {
	var arr = [10]int{10, 4, 12, 56, 90, 32, 12, 3, 54, 76}

	small := arr[0]
	var secondeSmall int

	for i := 1; i < len(arr); i++ {

		if small > arr[i] {

			secondeSmall = small
			small = arr[i]

		}

	}
	fmt.Println("The second smallest number is : ", secondeSmall)
}
