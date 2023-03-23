package main

import "fmt"

func main() {
	arr := [6]int{5, 8, 2, 9}
	value := 11
	num1, num2 := test2(arr, value)
	fmt.Printf("num1 : %d \nnum2 : %d\n", num1, num2)
}


func test2(arr [6]int, target int)(int,int){
	mySet := make(map[int]int)

	for i:=0; i<len(arr)-1; i++{
		match := target - arr[i]
		if value,ok := mySet[match]; ok{
			return value, i
		}else{
			mySet[arr[i]] = i
		}
	}
	return 0,0
}

func twoSum(arr []int, target int) (int,int) {
    mySet := make(map[int]int)
    for i:=0; i<len(arr); i++{
        match := target - arr[i]
        if value,ok :=mySet[match]; ok{
            return value,i
        }else{
            mySet[arr[i]]=i
        }
    }
    return 0,0
}
