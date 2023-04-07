package main

import "fmt"

func main() {
	arr := []int{5, 8, 2, 9}
	value := 11
	num1, num2 := twoSum(arr, value)
	fmt.Printf("num1 : %d \nnum2 : %d\n", num1, num2)
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
