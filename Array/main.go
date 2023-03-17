package main

import (
	"fmt"
)

func main() {
	var myArray [5]int

	for i := 0; i < len(myArray); i++ {
		fmt.Printf("Enter the elements %d :", i)
		fmt.Scan(&myArray[i])
	}
	fmt.Println(myArray)
}
