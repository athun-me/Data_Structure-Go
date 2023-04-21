package main

import (
	"fmt"
)

var heap []int

func heapSort(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from heap in descending order
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

// Heapify function to build max heap
func heapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l 
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}


//Heap sort diffrent method

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

func heapSort2(arr []int) {
	buildHeap(arr, len(arr)-1)
	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)
		buildHeap(arr, i-1)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	heapSort(arr)
	fmt.Println(arr)
}
