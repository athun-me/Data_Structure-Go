package main

import "fmt"

func quickSort(arr []int) []int {
	
	quickSortHelper(arr, 0, len(arr)-1)
	return arr
}

func quickSortHelper(arr []int, startIdx int, endIdx int) {
	if startIdx >= endIdx {
		return
	}

	pivotIdx := startIdx
	leftIdx := startIdx + 1
	rightIdx := endIdx

	for rightIdx >= leftIdx {

		if arr[leftIdx] > arr[pivotIdx] && arr[rightIdx] < arr[pivotIdx] {
			swap(arr, leftIdx, rightIdx)
			leftIdx++
			rightIdx--
		}
		if arr[leftIdx] <= arr[pivotIdx] {
			leftIdx++
		}
		if arr[rightIdx] >= arr[pivotIdx] {
			rightIdx--
		}

	}
	swap(arr, pivotIdx, rightIdx)
	quickSortHelper(arr, startIdx, rightIdx-1)
	quickSortHelper(arr, rightIdx+1, endIdx)
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	array := []int{1, 5, 10, 2, 6, 3}
	arr := quickSort(array)
	fmt.Println(arr)

}
