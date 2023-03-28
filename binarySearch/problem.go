package main

import "fmt"

func binaryIter(array [10]int, target int) int {
	starIdx := 0
	endIdx := len(array) - 1

	for starIdx <= endIdx {
		middleValue := starIdx + (endIdx-starIdx)/2
		if array[middleValue] == target {
			return middleValue
		} else if array[middleValue] < target {
			starIdx = middleValue + 1
		} else if array[middleValue] > target {
			endIdx = middleValue - 1
		}
	}
	return -1
}

func binaryRec(array [10]int, target int) int {
	return binaryRecHelper(array, target, 0, len(array)-1)
}

func binaryRecHelper(array [10]int, target int, startIdx int, endIdx int) int {

	middleValue := startIdx + (endIdx-startIdx)/2
	if startIdx > endIdx {
		return -1
	}
	
	if array[middleValue] == target {
		return middleValue
	}else if array[middleValue] < target {
		return binaryRecHelper(array, target, middleValue+1, endIdx)
	}else{
		return binaryRecHelper(array, target, startIdx, middleValue-1)
	}

}

func main() {
	arr := [10]int{10, 24, 32, 40, 52, 61, 78, 80, 99, 111}
	findIdex := binaryIter(arr, 52)
	if findIdex == -1 {
		fmt.Println("not found")
		return
	}
	fmt.Printf("the index is: %d, and the value is: %d\n", findIdex, arr[findIdex])
}
