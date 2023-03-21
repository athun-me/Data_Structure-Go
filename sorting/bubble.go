package main

import "fmt"

func bubbleSort(arr [10]int) {
	len := len(arr) - 1

	for len >= 0 {
		for i := 0; i < len; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		len--
	}
	fmt.Println(arr)
}

func main(){
	var arr = [10]int{10, 4, 12, 56, 90, 32, 12, 3, 54, 76}
	bubbleSort(arr)
}
