//Find the two numbers in the array that, when added, will equal 10.
package main

import (
	"fmt"
)

func findNumber(arr []int){
	len := len(arr)-1
	for i:=0; i<len-1; i++{
		for j:= i+1; j<len; j++{
			if arr[i]+arr[j] == 10 {
				fmt.Printf("num1 : %d, num2 :%d\n", arr[i], arr[j])
			}
		}
	}
}

func main(){
	var array = [10]int{1,2,3,4,5,6,7,8,9,10}
	findNumber(array[:])
}