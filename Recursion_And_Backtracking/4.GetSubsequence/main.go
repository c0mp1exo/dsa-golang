package main

import "fmt"

func main() {
	str := "abc"

	strlist := GetSubSequence(str, 0)
	for i, val := range strlist {
		fmt.Printf("value at idx %d is %s\n", i, val)
	}
}

func GetSubSequence(str string, idx int) []string {

	if idx == len(str) {
		return []string{""}
	}
	result := GetSubSequence(str, idx+1)
	ch := str[idx]
	var newResult []string
	for i := 0; i < len(result); i++ {
		newString := result[i]
		newResult = append(newResult, newString)
	}
	for i := 0; i < len(result); i++ {
		newString := string(ch) + result[i]
		newResult = append(newResult, newString)
	}
	return newResult

}
