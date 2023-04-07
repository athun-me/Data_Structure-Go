package main

import "fmt"

func mergSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middleIdx := len(arr) / 2
	firstHalf := arr[0:middleIdx]
	lastHalf := arr[middleIdx:]

	return join(mergSort(firstHalf), mergSort(lastHalf))
}

func join(firtHalf []int, lastHalf []int) []int {
	newArray := make([]int, len(firtHalf)+len(lastHalf))
	i, j, k := 0, 0, 0

	for i < len(firtHalf) && j < len(lastHalf) {
		if firtHalf[i] <= lastHalf[j] {
			newArray[k] = firtHalf[i]
			k++
			i++
		} else {
			newArray[k] = lastHalf[j]
			k++
			j++
		}
	}

	for i < len(firtHalf) {
		newArray[k] = firtHalf[i]
		k++
		i++
	}
	for j < len(lastHalf) {
		newArray[k] = lastHalf[j]
		k++
		j++
	}
	return newArray
}

func main() {
	var array = []int{4, 2, 10, 9, 43, 1, 5}
	arr := mergSort(array)
	fmt.Println(arr)
}
