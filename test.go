package main

import (
	"fmt"
)

var heap []int

func buildHeap(arr []int, l int) {
	heap = arr
	for i := parent(len(heap) - 1); i >= 0; i-- {
		shiftDown(i, l)
	}
}

func shiftDown(currentIdx int, endIdx int) {
	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {
		rightIdx := rightChild(currentIdx)

		var idxToShift int
		if rightIdx <= endIdx && heap[rightIdx] > heap[leftIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}

		if idxToShift <= endIdx && heap[idxToShift] > heap[currentIdx] {
			swap(heap, idxToShift, currentIdx)
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			return
		}
	}
}

func heapSort(arr []int) {
	buildHeap(arr, len(arr)-1)

	for i := len(arr) - 1; i >= 0; i-- {
		swap(arr, 0, i)
		buildHeap(arr, i-1)
	}
}

func peak() {
	fmt.Println(heap[0])
}

func shiftUp(currentIdx int) {
	parentIdx := parent(currentIdx)
	for currentIdx > 0 && heap[parentIdx] > heap[currentIdx] {
		swap(heap, currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = parent(currentIdx)
	}
}

func remove() {
	if len(heap) == 0 {
		fmt.Println("Heap is empty.")
		return
	}
	swap(heap, 0, len(heap)-1)
	heap = heap[:len(heap)-1]
	shiftDown(0, len(heap))
}

func insert(value int) {
	heap = append(heap, value)
	shiftUp(len(heap) - 1)
}

func display(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return (i * 2) + 1
}

func rightChild(i int) int {
	return (i * 2) + 2
}

func main() {
	arr := []int{4, 3, 2, 5, 1, 6, 7, 19, 3}

	heapSort(arr)
	// buildHeap(arr, len(arr)-1)
	fmt.Println(heap)
}
