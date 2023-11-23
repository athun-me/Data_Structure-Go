package main

import "fmt"

func insesrtSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
func bubbleSort(arr []int) {

	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
}

func main() {
	arr := []int{8, 8, 9, 4, 5, 6, 1, 7, 2, 3, 3}

	fmt.Println("Unsorted array:", arr)

	bubbleSort(arr)

	fmt.Println("Sorted array:", arr)
}
