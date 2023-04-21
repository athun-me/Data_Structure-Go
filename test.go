package main

import "fmt"

var heap []int

func leftChild(i int) int {
	return (2 * i) + 1
}
func rightChild(i int) int {
	return (2 * i) + 2
}
func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func shiftDown(currentIdx int, endIdx int) {

	leftIdx := leftChild(currentIdx)
	var idxToShift int

	for leftIdx <= endIdx {

		rightIdx := rightChild(currentIdx)

		if rightIdx <= endIdx && heap[rightIdx] > heap[leftIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}

		if heap[idxToShift] > heap[currentIdx] {
			swap(heap, currentIdx, idxToShift)
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			return
		}
	}
}

func buildHeap(arr []int, l int) {
	heap = arr
	for i := l; i >= 0; i-- {
		shiftDown(i, l)
	}
}

func heapSort(arr []int) {
	buildHeap(arr, len(arr)-1)
	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)
		buildHeap(arr, i-1)
	}
}

func main() {
	arr := []int{4, 3, 2, 1, 5, 6}
	
	heapSort(arr)
	fmt.Println(arr)
}
