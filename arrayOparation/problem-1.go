package main

import "fmt"

func main() {
	arr := [10]int{6, 2, 6, 5, 6, 8, 6, 3, 6, 10}
	value := 6
	sort(arr, value)
	
}

//Arranging all the value 6 at the tail

func sort(arr [10]int, value int){
	j := len(arr)-1
	
	for i:=0; i<len(arr); i++{
		if i ==j {
			break
		}
		if arr[i]==value{
			if arr[j] == value{
				j--
				i--
			}else{
				temp :=arr[i]
				arr[i]=arr[j]
				arr[j]=temp
			}
		}
	}
	fmt.Println(arr)
}

// The time complexity of this code is O(n), 
// where n is the length of the input array.
// This is because the code uses a single loop to iterate over the array once,
// and the time taken for each iteration is constant (i.e., O(1)). 
// Therefore, 
// the overall time complexity is proportional to the length of the array, which is O(n).

// The space complexity of this code is also O(1), which is constant. 
// This is because the code only uses a fixed amount of memory to store a few variables (e.g., i, j, temp), 
// and the amount of memory used does not depend on the size of the input array. 
// Therefore, 
// the space complexity is constant and does not grow as the size of the input increases.