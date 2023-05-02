package main

import "fmt"

func main() {
	arr := []int{1, 4, 4, 4, 5, 6, 7}
	length := len(arr)
	v := 4
	var j int

	for i := 0; i < length; i++ {

		if arr[i] == v {

			for j = i; j < length-1; j++ {
				arr[j] = arr[j+1]
			}
			i--
			length--
		}

	}

	arr = arr[:length]
	fmt.Println(arr)
}
