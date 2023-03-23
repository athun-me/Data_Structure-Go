package main

import "fmt"

func test(arr [4]int) {
	myset := make(map[int]int)

	for i := 0; i < len(arr)-1; i++ {
		count := 1
		for j := i + 1; j < len(arr); j++ {
			if _, ok := myset[arr[i]]; ok == false {
				if arr[i] == arr[j] {
					count++
				}
			}
		}
		myset[arr[i]] = count
	}
	fmt.Println(myset)
}

func test2(arr [7]int) {

	myset := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		myset[arr[i]]++
	}
	fmt.Println(myset)
}

func searchInsert(nums []int, target int) int {
    start := 0
    end := len(nums) - 1
    
    for start <= end {
        middle := (start + end) / 2
        
        if nums[middle] == target {
            return middle
        } else if nums[middle] > target {
            end = middle - 1
        } else {
            start = middle + 1
        }
    }
    
    return start
}


func main() {
	// var array = [4]int{1,3,5,6}
	// start := searchInsert(array[:],5)
	// fmt.Println(start)

	c:=r(2)
	fmt.Println(c)
}

// func (n int){
// 	var i,j,k, count = 0
// 	for i=n/2; i<=n; i++{
// 			for k=1; k<=n; k++{
// 				for j=1; j<=n; 2*j{
// 					count++
// 			}
// 		}
// 	}
// }

func r(n int)int{

	if n<=0 {
		return 1
	}
	return 1+ r(n-1)
}