package main

import "fmt"

func moveTwosToEnd(key int) {

	arr := [11]int{1, 2, 2, 3, 4, 5, 5, 6, 43, 435, 698}
	fmt.Println(arr)


	end := len(arr) - 1

	for i := 0; i <= end; {
		if arr[i] == key {
			temp := arr[i]
			arr[i] = arr[end]
			arr[end] = temp
			end--
		} else {
			i++
		}
	}

	
	fmt.Println(arr)
}

func main(){
	moveTwosToEnd(2)
}