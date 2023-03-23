package main

import (
	"fmt"
)

func countString(str string) {
	var result string
	count := 1
	for i := 0; i < len(str); i++ {
		if i == len(str)-1 || str[i] != str[i+1] {
			result += fmt.Sprintf("%d%c ", count, str[i])
			count = 1
		} else {
			count++
		}
	}
	fmt.Println(result)
}


func main() {
	countString("aaabbc")
}
