package main

import (
	"fmt"
)

func binarSearch(arr []int, target int) int {
	startIdx := 0
	endIdx := len(arr) - 1

	for startIdx <= endIdx {
		middleValue := startIdx + (endIdx-startIdx)/2
		if arr[middleValue] == target {
			return middleValue
		} else if arr[middleValue] < target {
			startIdx = middleValue + 1
		} else if arr[middleValue] > target {
			endIdx = middleValue - 1
		}
	}
	return -1
}

func binaryRec(arr[]int, target int)int{
	return binaryRecHelper(arr,target,0,len(arr)-1)
}
func binaryRecHelper(arr[]int, target int, startIdx int, endIdex int )int{
	middleValue := startIdx + (endIdex - startIdx)/2
	if startIdx > endIdex {
		return -1
	}
	if arr[middleValue] == target{
		return middleValue
	}else if arr[middleValue] < target{
		return binaryRecHelper(arr, target, middleValue+1, endIdex)
	}else{
		return binaryRecHelper(arr,target,startIdx, middleValue-1)
	}
}

func main() {
	var array = [10]int{10, 23, 35, 42, 53, 63, 70, 81, 99, 100}
	value := binaryRec(array[:], 42)
	fmt.Println(value)
}
