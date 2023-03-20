package main

import "fmt"

func binaryIter(array [10]int, target int)int{
	starIdx := 0
	endIdx := len(array)-1

	for starIdx <= endIdx {
		middleValue := starIdx + (endIdx - starIdx)/2
		if array[middleValue] == target{
			return middleValue
		}else if array[middleValue] < target{
			starIdx = middleValue + 1
		}else if array[middleValue] > target{
			endIdx = middleValue - 1
		}
	}
	return -1
}

func main() {
	arr := [10]int{10, 24, 32, 40, 52, 61, 78, 80, 99, 111}
	findIdex := binaryIter(arr, 32)
	fmt.Printf("the index is: %d, and the value is: %d\n", findIdex, arr[findIdex])
}
