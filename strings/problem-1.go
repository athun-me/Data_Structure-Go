package main

import "fmt"

func stringChange(str string, key int) string {
	newKey := key % 26
	charArray := []byte(str)

	for i := 0; i < len(str); i++ {
		letterPosition := str[i] + byte(newKey)
		if letterPosition <= 122 {
			charArray[i] = byte(letterPosition)
		} else {
			charArray[i] = byte(96 + letterPosition%122)
		}
	}
	return string(charArray)
}

func stringPalindrome(str string) {
	len := len(str) - 1
	var flage bool = true
	for i := 0; i < len; i++ {
		if str[i] != str[len] {
			flage = false
			break
		}
		len--
	}
	if flage == true {
		fmt.Println("string is palindrom")
	}else {
		fmt.Println("strin is not palidrom")
	}
}

func main() {
	stringPalindrome("malayalam")
}
