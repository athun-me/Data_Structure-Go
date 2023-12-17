package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 6, 2, 9, 1, 8, 7}
	arr2 := quickSort(arr)
	fmt.Println(arr2)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left []int
	var right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	return merge(mergeSort(arr[mid:]), mergeSort(arr[:mid]))
}
func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	var res []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		res = append(res, arr1[i:]...)
	}
	if j < len(arr2) {
		res = append(res, arr2[j:]...)
	}

	return res
}

func insertionSort(arr []int) {
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
	n := len(arr) - 1
	isTrue := true
	for isTrue {
		isTrue = false
		for i := 0; i < n; i++ {
			if arr[i] > arr[i+1] {
				isTrue = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		n--
	}
}

func sortLinear(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
