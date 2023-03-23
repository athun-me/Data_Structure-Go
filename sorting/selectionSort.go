package main

import "fmt"

func selectionSort(arr [10]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("Sorted array :", arr)
}

func main(){
	var array = [10]int{7, 6, 5, 1, 9, 8, 2, 4, 10, 3}
	selectionSort(array)
}