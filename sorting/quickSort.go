package main

import "fmt"

func quickSort(arr [10]int) [10]int {
	quickSortHelper(arr, 0, len(arr)-1)
	fmt.Println(arr)
	return arr

}

func quickSortHelper(arr [10]int, startIdx int, endIdx int) {
	if startIdx >= endIdx {
		return
	}
	pivotIdx := startIdx
	leftIdx := startIdx + 1
	rightIdx := startIdx

	for leftIdx <= rightIdx {
		if arr[leftIdx] > arr[pivotIdx] && arr[rightIdx] < arr[pivotIdx] {
			swap(arr, leftIdx, rightIdx)
			leftIdx++
			rightIdx--
		}
		if arr[leftIdx] <= arr[pivotIdx] {
			leftIdx++
		}
		if arr[leftIdx] >= arr[pivotIdx] {
			rightIdx--
		}
	}
	swap(arr, rightIdx, pivotIdx)
	quickSortHelper(arr, startIdx, rightIdx-1)
	quickSortHelper(arr, rightIdx+1, endIdx)
}

func swap(arr [10]int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	array := [10]int{6, 4, 8, 1, 5, 10, 7, 3, 9, 2}
	arr := quickSort(array)
	fmt.Println(arr)
}

