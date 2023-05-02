package main

import "fmt"

var heap []int

func leftChild(i int) int {
	return (i * 2) + 1
}
func rightChild(i int) int {
	return (i * 2) + 2
}

func parent(i int) int {
	return (i - 1) / 2
}
func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func buildHeap(arr []int, l int) {
	heap = arr

	for i := parent(len(heap) - 1); i >= 0; i-- {
		shiftDown(i, l)
	}
}

func shiftDown(currentIdx int, endIdx int) {
	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {
		var idxToShift int
		rightIdx := rightChild(currentIdx)

		if leftIdx <= endIdx && heap[leftIdx] < heap[rightIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}

		if heap[currentIdx] < heap[idxToShift] {
			swap(heap, currentIdx, idxToShift)
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			return
		}
	}
}

func heapSort(arr []int) {
	buildHeap(arr, len(arr)-1)

	for i := len(heap) - 1; i >= 0; i-- {
		swap(heap, 0, i)
		buildHeap(heap, i-1)
	}
}

func main() {
	array := []int{4, 2, 3, 8, 9, 8}

	heapSort(array)

	fmt.Println(heap)
	fmt.Println("Hello Afthab")
}
