package main

import "fmt"

func main() {
	arr := []int{4, 2, 8, 10, 2, 8, 28, 3}
	b := make([]bool, len(arr))
	
	for i := 0; i < len(arr); i++ {
		
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				b[i] = true
				b[j] = true
			}

		}
		if !b[i] {
			fmt.Printf("%d ",arr[i])
		}
	}
	fmt.Println()
}
