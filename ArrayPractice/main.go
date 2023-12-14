package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 8, 9, 2, 7, 2, 8}
	fmt.Println(mergSort(arr))
}

func quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var left []int
	var right []int
	pivot := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(quick(left), pivot), quick(right)...)
}

func bubble(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
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

func mergSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2

	x := mergSort(arr[:mid])
	y := mergSort(arr[mid:])
	return merg(x, y)
}

func merg(arr, arr2 []int) []int {
	i, j := 0, 0

	var res []int

	for i < len(arr) && j < len(arr2) {
		if arr[i] < arr2[j] {
			res = append(res, arr[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	if i < len(arr) {
		res = append(res, arr[i:]...)
	}
	if j < len(arr2) {
		res = append(res, arr2[j:]...)
	}
	return res
}
