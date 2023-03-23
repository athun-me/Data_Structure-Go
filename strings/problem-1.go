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

func main() {
	str := stringChange("hai", 2)
	fmt.Println(str)
}
