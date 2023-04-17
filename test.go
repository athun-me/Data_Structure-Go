package main

import "fmt"

func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, i int) {
	large := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] < arr[large] {
		large = l
	}
	if r < n && arr[r] < arr[large] {
		large = r
	}
	if large != i {
		arr[i], arr[large] = arr[large], arr[i]
		heapify(arr, n, large)
	}
}

func main() {
	arr := []int{2, 3, 4, 2, 56, 133}
	heapSort(arr)
	fmt.Println(arr)
}
