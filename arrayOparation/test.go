package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	res1 := mergeSort(arr[len(arr)/2:])
	res2 := mergeSort(arr[:len(arr)/2])

	return merge(res1, res2)
}

func merge(arr1, arr2 []int) []int {
	res := []int{}

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			res = append(res, arr2[j])
			j++
		} else {
			res = append(res, arr1[i])
			i++
		}
	}
	for i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}

	for j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}
	return res
}
func inserSort(arr []int) {
	for i := 1; i < len(arr); i++ {
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
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{9, 8, 5, 6, 2, 3, 1, 7, 4}
	bubbleSort(arr)

	fmt.Println(arr)
}
