package main

import "fmt"

func insertionSort(arr [10]int) {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
	fmt.Println(arr)
}

func main() {
	var array = [10]int{6, 4, 8, 1, 5, 10, 7, 3, 9, 2}
	insertionSort(array)
}
